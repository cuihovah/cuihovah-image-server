package image

import (
    "github.com/nfnt/resize"
    "image"
    _ "image/gif"
    _ "image/jpeg"
    "image/png"
    "math"
    "os"
)

const DEFAULT_MAX_WIDTH float64 = 160
const DEFAULT_MAX_HEIGHT float64 = 120

// 计算图片缩放后的尺寸
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
    ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
    return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// 生成缩略图
func MakeThumbnail(imagePath, savePath string) (error) {
    
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        return err
    }

    b := img.Bounds()
    width := b.Max.X
    height := b.Max.Y

    w, h := calculateRatioFit(width, height)

    // 调用resize库进行图片缩放
    m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

    // 需要保存的文件
    imgfile, _ := os.Create(savePath)
    defer imgfile.Close()

    // 以PNG格式保存文件
    err = png.Encode(imgfile, m)
    if err != nil {
        return err
    }

    return nil
}