package log

import "k8s.io/klog/v2"

func Codec() klog.Verbose {
	return klog.V(7)
}

func Network() klog.Verbose {
	return klog.V(8)
}
