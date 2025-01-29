# sudo apt install notify-tools
inotifywait --monitor --event modify *.go |
while read file event; do
    clear # clear the console
    go run . # run the code
    # go build -o output/csv_to_db
    # /output/csv_to_db 
done
