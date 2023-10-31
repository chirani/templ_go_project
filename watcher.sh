watch_folder="./"

# The command to run
command_to_run_templ="templ generate"
# Monitor the folder for changes
inotifywait -m -r -e modify,move,create,delete "$watch_folder" |
while read -r directory event file
do
  if [[ "$file" == *.templ ]]; then
    echo "Detected a change in a .templ file: $file"
    echo "Running $command_to_run_templ..."
    $command_to_run_templ
  fi
done