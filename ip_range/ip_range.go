package iprange

import (
	"fmt"
	"net"
	"strings"

	"github.com/charlienet/go-mixed/bytesconv"
)

type IpRange struct {
	segments []ipSegment
}

type ipSegment interface {
	Contains(net.IP) bool
}

type singleIp struct {
	ip net.IP
}

func (i *singleIp) Contains(ip net.IP) bool {
	return i.ip.Equal(ip)
}

type cidrSegments struct {
	cidr *net.IPNet
}

func (i *cidrSegments) Contains(ip net.IP) bool {
	return i.cidr.Contains(ip)
}

type rangeSegment struct {
	start rangeIP
	end   rangeIP
}

type rangeIP struct {
	Hight uint64
	Lower uint64
}

func (r *rangeSegment) Contains(ip net.IP) bool {
	ih, _ := bytesconv.BigEndian.BytesToUInt64(ip[:8])
	i, _ := bytesconv.BigEndian.BytesToUInt64(ip[8:])

	return ih >= r.start.Hight && ih <= r.end.Hight && i >= r.start.Lower && i <= r.end.Lower
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
	nip := net.ParseIP(ip)
	if nip == nil {
		return false
	}

	for _, v := range r.segments {
		if v.Contains(nip) {
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

		start := net.ParseIP(ips[0])
		end := net.ParseIP(ips[1])
		if start == nil {
			return nil, fmt.Errorf("IP范围起始地址格式错误:%s", ips[0])
		}

		if end == nil {
			return nil, fmt.Errorf("IP范围结束地址格式错误:%s", ips[0])
		}

		sh, _ := bytesconv.BigEndian.BytesToUInt64(start[:8])
		s, _ := bytesconv.BigEndian.BytesToUInt64(start[8:])
		eh, _ := bytesconv.BigEndian.BytesToUInt64(end[:8])
		e, _ := bytesconv.BigEndian.BytesToUInt64(end[8:])

		return &rangeSegment{start: rangeIP{
			Hight: sh, Lower: s},
			end: rangeIP{Hight: eh, Lower: e}}, nil

	case strings.Contains(ip, "/"):
		if _, cidr, err := net.ParseCIDR(ip); err != nil {
			return nil, err
		} else {
			return &cidrSegments{cidr: cidr}, nil
		}
	default:
		i := net.ParseIP(ip)
		if i == nil {
			return nil, fmt.Errorf("格式错误, 不是有效的IP地址:%s", ip)
		}

		return &singleIp{ip: i}, nil
	}
}
