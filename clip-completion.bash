function _clip_complete_()
{
    local cmd="${1##*/}"
    local word=${COMP_WORDS[COMP_CWORD]}
    local line=${COMP_LINE}
    COMPREPLY=($($cmd -c $word $line))
}

complete -F _clip_complete_ clip
