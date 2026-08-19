package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestGenerateShortLink(t *testing.T) {
	initiallink1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortlink1 := GenerateShortLink(initiallink1, UserId)

	initiallink2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortlink2 := GenerateShortLink(initiallink2, UserId)

	initiallink3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortlink3 := GenerateShortLink(initiallink3, UserId)

	assert.Equal(t, shortlink1, "jTa4L57P")
	assert.Equal(t, shortlink2, "d66yfx7N")
	assert.Equal(t, shortlink3, "dhZTayYQ")
}
