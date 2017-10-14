package routes

import (
    "image"
    "image/png"
    "image/color"
    "math/rand"
    "github.com/gin-gonic/gin"
)

func randomImage(c *gin.Context) {
    // From: http://41j.com/blog/2015/03/creating-a-png-image-in-golang/

    myimage := image.NewRGBA(image.Rectangle{image.Point{0,0},image.Point{200,200}})

    // This loop just fills the image with random data
    for x := 0; x < 200; x++ {
        for y := 0; y < 200; y++ {
            c := color.RGBA{uint8(rand.Intn(255)),uint8(rand.Intn(255)),uint8(rand.Intn(255)),255}
            myimage.Set(x,y,c)
        }
    }

    c.Writer.Header().Set("Content-Type", "image/png")

    png.Encode(c.Writer, myimage)

    c.Writer.Flush()
}
