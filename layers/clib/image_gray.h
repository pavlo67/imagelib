#include <inttypes.h>

typedef struct {
    int32_t x_width, y_height;
    int32_t x_min, y_min;
    int32_t stride;
    unsigned char* pixels;
} ImageGray;

typedef struct {
    ImageGray image_gray;
    int32_t*  classes;
    int8_t    classesNum;
} Layer;

ImageGray CheckConnect (ImageGray);
Layer     BrClasses    (Layer);
