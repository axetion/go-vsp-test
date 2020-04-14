package main

/*
// Stubs for the virtuals we're not going to override.
void dummy(void) {}		// return nothing
int dummy2(void) { return 0; }		// return PLUGIN_CONTINUE

// Our exported Go functions.
extern unsigned char Load(void *this, void *fac1, void *fac2);
extern void Unload(void *this);
extern const char *Description(void *this);
extern void PutInServer(void *this, void *player, const char *name);

void *vtable[20] = {
	&Load,			// ::Load
	&Unload,		// ::Unload
	&dummy,			// ::Pause
	&dummy,			// ::UnPause
	&Description,	// ::GetPluginDescription
	&dummy,			// ::LevelInit
	&dummy,			// ::ServerActivate
	&dummy,			// ::GameFrame
	&dummy,			// ::LevelShutdown
	&dummy,			// ::ClientActive
	&dummy,			// ::ClientDisconnect
	&PutInServer,	// ::ClientPutInServer
	&dummy,			// ::SetCommandClient
	&dummy,			// ::ClientSettingsChanged
	&dummy2,		// ::ClientConnect
	&dummy2,		// ::ClientCommand
	&dummy2,		// ::NetworkIDValidated
	&dummy,			// ::OnQueryCvarValueFinished
	&dummy,			// ::OnEdictAllocated
	&dummy			// ::OnEdictFreed
};

void *vptr = &vtable;
*/
import "C"

var vptr = &C.vptr
