# Define the array of options
options=("init" "commit" "add")

# Function to handle autocomplete
_commitix_autocomplete() {
    local cur="${COMP_WORDS[COMP_CWORD]}"
    COMPREPLY=( $(compgen -W "${options[*]}" -- "$cur") )
}

# Register the autocomplete function for the 'commitix' command
complete -F _commitix_autocomplete commitix
