APP       := TestApp
TARGET    := TestAppServer
MFLAGS    :=
DFLAGS    :=
CONFIG    := config.conf
STRIP_FLAG:= N
J2GO_FLAG:= 

libpath=${subst :, ,$(GOPATH)}
$(foreach path,$(libpath),$(eval -include $(path)/src/github.com/TarsCloud/TarsGo/tars/makefile.tars))
