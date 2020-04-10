#pragma once

#ifndef QAMEL_HANDLEMSG_H
#define QAMEL_HANDLEMSG_H

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus

// Class
class HandleMsg;

extern "C" {
#endif

// Properties

// Signals
void HandleMsg_SendDanMu(void* ptr, char* p0);
void HandleMsg_SendGift(void* ptr, char* p0);
void HandleMsg_SendWelCome(void* ptr, char* p0);
void HandleMsg_SendWelComeGuard(void* ptr, char* p0);
void HandleMsg_SendGreatSailing(void* ptr, char* p0);
void HandleMsg_SendOnlineChanged(void* ptr, int p0);
void HandleMsg_SendFansChanged(void* ptr, int p0);
void HandleMsg_SendMusicURI(void* ptr, char* p0, char* p1, char* p2);

// Register
void HandleMsg_RegisterQML(char* uri, int versionMajor, int versionMinor, char* qmlName);

#ifdef __cplusplus
}
#endif

#endif
