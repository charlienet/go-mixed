package iprange

import (
	"fmt"
	"net/netip"
	"regexp"
	"strconv"
	"strings"
)

var maskPattern = regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)

type IpRange struct {
	segments []ipSegment
}

type ipSegment interface {
	Contains(netip.Addr) bool
}

type singleIp struct {
	ip netip.Addr
}

func (i *singleIp) Contains(ip netip.Addr) bool {
	return i.ip.Compare(ip) == 0
}

type prefixSegments struct {
	prefix netip.Prefix
}

func (i *prefixSegments) Contains(ip netip.Addr) bool {
	return i.prefix.Contains(ip)
}

type rangeSegment struct {
	start netip.Addr
	end   netip.Addr
}

func (r *rangeSegment) Contains(ip netip.Addr) bool {
	return ip.Compare(r.start) >= 0 && ip.Compare(r.end) <= 0
}

// IP范围判断，支持以下规则:
// 单IP地址，如 192.168.100.2
// IP范围, 如 192.168.100.120-192.168.100.150
// 掩码模式，如 192.168.2.0/24
func NewRange(ip ...string) (*IpRange, error) {
	seg := make([]ipSegment, 0, len(ip))

	for _, i := range ip {
		if s, err := createSegment(i); err != nil {
			return nil, err
		} else {
			seg = append(seg, s)
		}
	}

	return &IpRange{segments: seg}, nil
}

func (r *IpRange) Contains(ip string) bool {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return false
	}

	for _, v := range r.segments {
		if v.Contains(addr) {
			return true
		}
	}

	return false
}

func createSegment(ip string) (ipSegment, error) {
	switch {
	case strings.Contains(ip, "-"):
		ips := strings.Split(ip, "-")
		if len(ips) != 2 {
			return nil, fmt.Errorf("IP范围定义错误:%s", ip)
		}

		start, err := netip.ParseAddr(ips[0])
		if err != nil {
			return nil, err
		}

		end, err := netip.ParseAddr(ips[1])
		if err != nil {
			return nil, err
		}

		return &rangeSegment{
			start: start,
			end:   end,
		}, nil

	case strings.Contains(ip, "/"):
		sec := strings.Split(ip, "/")
		ip := sec[0]
		mask := sec[1]

		if maskPattern.MatchString(mask) {
			mask = strconv.Itoa(MaskToBits(mask))
		}

		if prefix, err := netip.ParsePrefix(ip + "/" + mask); err != nil {
			return nil, err
		} else {
			return &prefixSegments{prefix: prefix}, nil
		}
	default:
		i, err := netip.ParseAddr(ip)
		if err != nil {
			return nil, fmt.Errorf("格式错误, 不是有效的IP地址:%s", ip)
		}

		return &singleIp{ip: i}, nil
	}
}
