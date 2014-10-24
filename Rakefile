task :build do
  system("go build -o httpstream httpstream.go")
end

task :run do
  system("bundle exec thin start -R config.ru")
end