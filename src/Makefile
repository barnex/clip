include $(GOROOT)/src/Make.inc

all: clip

clip: pkg/_obj
	$(GC) -Ipkg/_obj main.go
	$(LD) -Lpkg/_obj -o $@ main.$O

pkg/_obj: pkg/*.go
	gomake -C pkg

.PHONY:	clean
clean:
	rm -f clip *.$O
	gomake clean -C pkg

