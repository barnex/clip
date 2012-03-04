function _clip_complete_()
{
    #local cmd="${1##*/}"
    local cmd=$COMP_WORDS
    #echo COMP_WORDS $COMP_WORDS
    #echo COMP_CWORD $COMP_CWORD
    #echo COMP_LINE $COMP_LINE 
	COMPREPLY=($($cmd complete $COMP_CWORD $COMP_LINE))
}

complete -F _clip_complete_ clip
