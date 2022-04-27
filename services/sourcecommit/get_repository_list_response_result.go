/*
 * api
 *
 * <br/>https://sourcecommit.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sourcecommit

type GetRepositoryListResponseResult struct {

	// total repository count
	Total *int32 `json:"total,omitempty"`

	Repository []*GetRepositoryListResponseResultRepository `json:"repository,omitempty"`
}
