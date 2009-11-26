include $(GOROOT)/src/Make.$(GOARCH)

TARG=kvs
GOFILES=\
	kvs.go\

include $(GOROOT)/src/Make.pkg
