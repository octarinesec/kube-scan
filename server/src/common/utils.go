package common

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/toolkits/slice"
	networkingv1 "k8s.io/api/networking/v1"
	"sort"
	"strings"
)

func EqualsMaps(m1 map[string]string, m2 map[string]string) bool {
	if len(m1) != len(m2) || len(m1) == 0 {
		return false
	}

	for k, v := range m1 {
		if v2, ok := m2[k]; !ok || v != v2 {
			return false
		}
	}

	return true
}

func MapAsString(m map[string]string) string {
	arr := make([]string, 0)
	for k, v := range m {
		arr = append(arr, fmt.Sprintf("%v: %v", k, v))
	}
	return strings.Join(arr, ", ")
}

func LabelsAsString(labels map[string]string) string {
	if labels == nil {
		return ""
	}

	labelsList := make([]string, 0)
	for k, v := range labels {
		labelsList = append(labelsList, fmt.Sprintf("%v: %v", k, v))
	}

	return strings.Join(labelsList, ",")
}

func JoinPolicyTypes(types []networkingv1.PolicyType, seperator string) string {
	arr := make([]string, len(types))
	for i, t := range types {
		arr[i] = string(t)
	}

	return strings.Join(arr, seperator)
}

func StringsDiff(s1 []string, s2 []string) []string {
	res := make([]string, 0)
	for _, s := range s1 {
		if !slice.ContainsString(s2, s) {
			res = append(res, s)
		}
	}
	return res
}

func SortAndHash(hashMap map[string]interface{}) ([]string, string) {
	var keys []string
	for k := range hashMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buffer bytes.Buffer
	for _, s := range keys {
		buffer.WriteString(fmt.Sprintf("%v", hashMap[s]))
	}

	hash := fmt.Sprintf("%x", md5.Sum(buffer.Bytes()))
	return keys, hash
}

func HasPrefix(s string, prefix ...string) bool {
	for _, p := range prefix {
		if strings.HasPrefix(s, p) {
			return true
		}
	}

	return false
}
