#!/bin/sh

print_help() {
	printf "Description:\n\tManipulates tokens generated with uuidgen.\nSynopsis:\n[OPTIONS] file\n\n\t[-n|--new DAYS]:\tcreates new token with expiration in <DAYS>\n\t[-r|--remove TOKEN]:\tremoves specified <TOKEN>\n\t[-i|--info TOKEN]:\tprints expiration date of specified <TOKEN>\n\t[-h|--help]:\t\tprints this help\nRequirements:\n\tuuidgen\n"
}


# $1=expiration $2=file
create_token() {
	token="$( uuidgen )"
	expiration="$( date -d "+$1 minutes" '+%Y-%m-%dT%H:%M:%SZ' )"
	echo "$token,$expiration" >>"$2"
	echo "$token"
}

# $1=token $2=file
remove_token() {
	sed -i "/^$1,.*$/d" "$2"
}

# $1=token $2=file
token_expiration_info() {
	grep "^$1,.*$" $2 | head -n1 | cut -d, -f2
}

eval set -- "$( getopt -o "n:r:i:h" -l "new:,remove:,info:,help" -- "$@" )"

expiration_days=""
token2remove=""
token2info=""

while [ $# -gt 0 ]; do
	case $1 in
		-n|--new)
			expiration_days=$2
			shift
			;;
		-r|--remove)
			token2remove="$2"
			shift
			;;
		-i|--info)
			token2info="$2"
			shift
			;;
		-h|--help)
			print_help
			;;
		--)
			shift
			break
			;;
		*)
			echo "Unknown option $1" >&2
			exit 1
			;;
	esac
	shift
done

if [ $# -ne 1 ]; then
	echo "Required exactly 1 argument, $# was given."
	exit 1
fi

tf="$1"

[ -f "$tf" ] || touch "$tf"

[ -n "$token2info" ] && token_expiration_info "$token2info" "$tf"
[ -n "$token2remove" ] && remove_token "$token2remove" "$tf"
if [ -n "$( echo "$expiration_days" | grep '^[0-9]\+$' )" ]; then
	create_token "$expiration_days" "$tf"
elif [ -n "$expiration_days" ]; then
	 echo "Expiration days specified after new paramater must be number, $expiration_days was given." >&2
	 exit 1
fi

exit 0
