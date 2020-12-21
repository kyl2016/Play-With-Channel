package client

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./lib -lcompareApi

#include <stdlib.h>
#include "compare.h"
#include "types.h"

int wrapLoad(void* handle, const int id, IVS_DATA_INFO* dataInfo, IVS_DEVICE_INFO* deviceInfo, unsigned char* data, int capacity) {
	dataInfo->data = data;
	dataInfo->deviceInfo = deviceInfo;
	return load(handle, id, dataInfo, capacity);
}
*/

func (c *gpu) Load(id int, features [][]byte, capacity int) error {
	if code := C.wrapLoad(
		c.handle,
		C.int(id),
		&C.IVS_DATA_INFO{
			dataType: C.FLOAT32,
			featureLength: C.int(c.featureLength),
			dataLength: C.uint(len(data)),
		},
		c.device,
		(*C.uchar)(dataPtr),
		C.int(c.featureLength*capacity*4))

	)
}