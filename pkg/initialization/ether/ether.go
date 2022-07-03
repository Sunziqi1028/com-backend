/**
 * @Author: Sun
 * @Description:
 * @File:  bounty
 * @Version: 1.0.0
 * @Date: 2022/7/1 20:38
 */

package ether

import (
	"ceres/pkg/service/ether"
)

func Init() error {
	ether.GetAllContractAddresses()
	return nil
}
