//go:generate mockgen -source ../domain/blogdb.go -destination=mock/blogdb.gen.go -package mock
//go:generate mockgen -source ../domain/fileutil.go -destination=mock/fileutil.gen.go -package mock
//go:generate mockgen -source context.go -destination=mock/context.gen.go -package mock
package app
