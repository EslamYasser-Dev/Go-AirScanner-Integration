// twain_wrapper.cpp
#include "twain-dsm.h"
#include "twain-wrapper.h"
#include <windows.h>
#include <stdio.h>
#include <string>

// Global TWAIN interfaces
TW_IDENTITY g_appID = {0};
TW_IDENTITY g_srcID = {0};
TW_HANDLE g_mem = nullptr;
TW_USERINTERFACE g_ui = {0};
TW_CAPABILITY g_cap = {0};
bool g_bEnabled = false;

// Forward declarations
TW_UINT16 DSM_Entry(TW_IDENTITY*, TW_IDENTITY*, TW_UINT32, TW_UINT16, TW_UINT32, TW_UINT32);
TW_UINT16 DSM_CloseDSMS(TW_IDENTITY*);

int twain_init() {
    g_appID.Id = 0;
    g_appID.Version.MajorNum = 1;
    g_appID.Version.MinorNum = 0;
    g_appID.Version.Language = TWLG_USA;
    g_appID.Version.Country = TWCY_USA;
    strcpy_s(g_appID.Version.Info, "Go TWAIN Client");
    g_appID.ProtocolMajor = TWON_PROTOCOLMAJOR;
    g_appID.ProtocolMinor = TWON_PROTOCOLMINOR;
    g_appID.SupportedGroups = DG_IMAGE | DG_CONTROL;
    strcpy_s(g_appID.Manufacturer, "Go Scanner");
    strcpy_s(g_appID.ProductFamily, "Scanners");
    strcpy_s(g_appID.ProductName, "Go TWAIN");

    TW_UINT16 rc = DSM_Entry(&g_appID, nullptr, DG_CONTROL, DAT_PARENT, MSG_OPENDSM, (TW_UINT32)GetForegroundWindow());
    if (rc != TWRC_SUCCESS) {
        return 0;
    }
    return 1;
}

int twain_select_source() {
    TW_UINT16 rc = DSM_Entry(&g_appID, &g_srcID, DG_CONTROL, DAT_IDENTITY, MSG_USERSELECT, 0);
    return (rc == TWRC_SUCCESS) ? 1 : 0;
}

int twain_open_source() {
    TW_UINT16 rc = DSM_Entry(&g_appID, &g_srcID, DG_CONTROL, DAT_SESSION, MSG_OPENDS, 0);
    if (rc == TWRC_SUCCESS) {
        g_bEnabled = false;
        return 1;
    }
    return 0;
}

void twain_close_source() {
    if (g_bEnabled) {
        DSM_Entry(&g_appID, &g_srcID, DG_CONTROL, DAT_USERINTERFACE, MSG_DISABLEDS, (TW_UINT32)&g_ui);
        g_bEnabled = false;
    }
    DSM_Entry(&g_appID, &g_srcID, DG_CONTROL, DAT_SESSION, MSG_CLOSEDS, 0);
}

int twain_enable_adf(int enable) {
    if (!g_bEnabled) {
        g_ui.hParent = GetForegroundWindow();
        g_ui.ShowUI = FALSE;
        g_ui.ModalUI = TRUE;
        DSM_Entry(&g_appID, &g_srcID, DG_CONTROL, DAT_USERINTERFACE, MSG_ENABLEDS, (TW_UINT32)&g_ui);
        g_bEnabled = true;
    }

    g_cap.Cap = CAP_FEEDERENABLED;
    g_cap.ConType = TWON_ONEVALUE;
    TW_UINT16 rc = DSM_Entry(&g_appID, &g_srcID, DG_CONTROL, DAT_CAPABILITY, MSG_SET, (TW_UINT32)&g_cap);
    return (rc == TWRC_SUCCESS) ? 1 : 0;
}

int twain_set_dpi(int dpi) {
    g_cap.Cap = ICAP_XRESOLUTION;
    g_cap.ConType = TWON_ONEVALUE;
    g_cap.hContainer = GlobalAlloc(GMEM_MOVEABLE, sizeof(TW_ONEVALUE));
    if (!g_cap.hContainer) return 0;

    TW_ONEVALUE* pOne = (TW_ONEVALUE*)GlobalLock(g_cap.hContainer);
    pOne->ItemType = TWTY_FIX32;
    pOne->Item = TW_FIX32FromFloat((float)dpi);
    GlobalUnlock(g_cap.hContainer);

    TW_UINT16 rc = DSM_Entry(&g_appID, &g_srcID, DG_CONTROL, DAT_CAPABILITY, MSG_SET, (TW_UINT32)&g_cap);
    GlobalFree(g_cap.hContainer);
    return (rc == TWRC_SUCCESS) ? 1 : 0;
}

int twain_set_color(int color_mode) {
    g_cap.Cap = ICAP_PIXELTYPE;
    g_cap.ConType = TWON_ONEVALUE;
    g_cap.hContainer = GlobalAlloc(GMEM_MOVEABLE, sizeof(TW_ONEVALUE));
    if (!g_cap.hContainer) return 0;

    TW_ONEVALUE* pOne = (TW_ONEVALUE*)GlobalLock(g_cap.hContainer);
    pOne->ItemType = TWTY_UINT16;
    pOne->Item = (color_mode == 1) ? TWPT_RGB : TWPT_GRAY;
    GlobalUnlock(g_cap.hContainer);

    TW_UINT16 rc = DSM_Entry(&g_appID, &g_srcID, DG_CONTROL, DAT_CAPABILITY, MSG_SET, (TW_UINT32)&g_cap);
    GlobalFree(g_cap.hContainer);
    return (rc == TWRC_SUCCESS) ? 1 : 0;
}

int twain_acquire(const char* output_dir) {
    // Full acquisition loop requires MSG_XFERREADY, file transfer, etc.
    // This is a simplified placeholder.
    // You'd need to implement the full state machine.
    return -1; // Not implemented in this snippet
}

void twain_exit() {
    if (g_bEnabled) {
        DSM_Entry(&g_appID, &g_srcID, DG_CONTROL, DAT_USERINTERFACE, MSG_DISABLEDS, (TW_UINT32)&g_ui);
        g_bEnabled = false;
    }
    DSM_Entry(&g_appID, nullptr, DG_CONTROL, DAT_PARENT, MSG_CLOSEDSMS, 0);
}