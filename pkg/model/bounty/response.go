/**
 * @Author: Sun
 * @Description:
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2022/6/29 13:17
 */

package bounty

type GetStartupsResponse struct {
	Startups map[uint64]string
}

type CreateBountyResponse struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}
