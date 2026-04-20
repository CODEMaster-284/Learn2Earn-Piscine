INTERVIEW=$(grep -h "SEE INTERVIEW" streetscd ./* | grep -oE '[0-9]+')
echo "$INTERVIEW"
cat interviews/interview-"$INTERVIEW"
echo "$MAIN_SUSPECT"
echo ""