#pragma once

#ifndef QAMEL_CONNECTFEEDBACK_H
#define QAMEL_CONNECTFEEDBACK_H

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus

// Class
class ConnectFeedBack;

extern "C" {
#endif

// Properties

// Signals
void ConnectFeedBack_SendFansNums(void* ptr, int p0);
void ConnectFeedBack_SendCompInfo(void* ptr, char* p0);
void ConnectFeedBack_SendErr(void* ptr, int p0);

// Register
void ConnectFeedBack_RegisterQML(char* uri, int versionMajor, int versionMinor, char* qmlName);

#ifdef __cplusplus
}
#endif

#endif
