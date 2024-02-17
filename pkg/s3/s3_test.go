package s3

import (
	"crypto/rand"
)

const testFilename = "test_filename"

func (s *Suite) TestStore() {
	s.Run("success", func() {
		bytes := make([]byte, 128)
		_, err := rand.Read(bytes)
		s.Require().NoError(err)

		err = s.client.Store(s.ctx, testFilename, bytes)
		s.Require().NoError(err)
	})
}
