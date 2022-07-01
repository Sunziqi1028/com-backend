/**
 * @Author: Sun
 * @Description:
 * @File:  bounty
 * @Version: 1.0.0
 * @Date: 2022/7/1 20:38
 */

package bounty

import "ceres/pkg/service/bounty"

func Init() error {
	bounty.GetAllContractAddresses()
	return nil
}
