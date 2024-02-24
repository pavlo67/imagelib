#include <stdio.h>
#include "image_gray.h"

ImageGray CheckConnect(ImageGray img) {
    printf("x_width: %i, y_height: %i, x_min: %i, y_min: %i, stride: %i, pix[0]: %i, pix[1]: %i, pix[2]: %i\n",
        img.x_width, img.y_height, img.x_min, img.y_min, img.stride, img.pix[0], img.pix[1], img.pixels[2]
    );

    return img;
}

Layer BrClasses(Layer lyr) {
    //    printf("x_width: %i, y_height: %i, x_min: %i, y_min: %i, stride: %i, pix[0]: %i, pix[1]: %i, pix[2]: %i\n",
    //        img.x_width, img.y_height, img.x_min, img.y_min, img.stride, img.pix[0], img.pix[1], img.pixels[2]
    //    );

    return lyr;
}
