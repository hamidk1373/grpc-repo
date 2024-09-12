s=nil
c=nil

gen:
	protoc --go_out=. --go-grpc_out=. allprotos/$(s).proto

gencommon:
	protoc --go_out=. --go-grpc_out=. allprotos/common.proto
	rm -rf commonpb
	mkdir commonpb
	mv gitlab.com/shanpanze/grpc-repo/commonpb/* ./commonpb
	rm -rf gitlab.com

pub: 
	git add .
	git commit -am "$(c)"
	git push
	./publish	