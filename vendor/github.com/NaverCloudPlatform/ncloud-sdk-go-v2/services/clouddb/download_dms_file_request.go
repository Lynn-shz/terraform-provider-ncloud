/*
 * clouddb
 *
 * Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2
 *
 * API version: 2018-11-01T03:57:11Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddb

type DownloadDmsFileRequest struct {

	// 클라우드DB인스턴스번호
CloudDBInstanceNo *string `json:"cloudDBInstanceNo"`

	// 파일이름
FileName *string `json:"fileName"`
}
