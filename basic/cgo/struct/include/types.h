
typedef struct DataInfo {
    IVS_DATA_TYPE dataType;
    int featureLength;
    IVS_DEVICE_INFO* deviceInfo;
    unsigned char* data;
    unsigned int dataLength;
}IVS_DATA_INFO;

typedef enum DataType{
    UINT8 = 0;
    FLOAT32 = 1;
}IVS_DATA_TYPE;

#define DEVICE_NAME_LENGTH 64
typedef struct DeviceInfo {
    int deviceNo;
    char deviceName[DEVICE_NAME_LENGTH];
}IVS_DEVICE_INFO;