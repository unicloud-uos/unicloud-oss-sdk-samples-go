package s3lib

import (
	"os"

	"github.com/journeymidnight/yig/api/datatype"
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/service/s3"
)

func (s3client *S3Client) GenTestObjectUrl(bucketName, objectKey string) string {
	return "http://" + *s3client.Client.Config.Endpoint + string(os.PathSeparator) + bucketName + string(os.PathSeparator) + objectKey
}

func TransferToS3AccessControlPolicy(policy *datatype.AccessControlPolicy) (s3policy *s3.AccessControlPolicy) {
	s3policy = new(s3.AccessControlPolicy)
	s3policy.Owner = new(s3.Owner)
	s3policy.Owner.ID = aws.String(policy.ID)
	s3policy.Owner.DisplayName = aws.String(policy.DisplayName)

	for _, p := range policy.AccessControlList {
		grant := new(s3.Grant)
		grant.Grantee = new(s3.Grantee)
		grant.Grantee.ID = aws.String(p.Grantee.ID)
		grant.Grantee.DisplayName = aws.String(p.Grantee.DisplayName)
		grant.Grantee.URI = aws.String(p.Grantee.URI)
		grant.Grantee.Type = aws.String(p.Grantee.XsiType)
		//	grant.Grantee.EmailAddress = aws.String(p.Grantee.EmailAddress)
		grant.Permission = aws.String(p.Permission)
		s3policy.Grants = append(s3policy.Grants, grant)
	}
	return
}

// Generate 5M part data
func GenMinimalPart() []byte {
	return RandBytes(5 << 20)
}
