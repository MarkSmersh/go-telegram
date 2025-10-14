<b>rm</b> [<u>options</u>] [<u>reply</u>]

<code>-s</code> <u>messageid</u> 
Message ID to start from

<code>-e</code> <u>messageid</u> 
Message ID to end with. Must be higher than <code>-s</code>

<code>-t</code> <u>number</u> 
From 1 to 100. Deletes messages in the range from original Message ID to Message ID - <u>number</u>  

<b>EXAMPLES</b>
<b>rm</b> <u>reply</u>
Deletes replied messsage

<b>rm -t 10</b>
Deletes the original message and existing messsages from above. Previous messages' ids should be in range of 100 regarding to the original message, otherwise those message will be skipped.

<b>rm -s 100 -e 200</b>
Deletes messages with ids in range of 100 to 200.

<b>rm -s 50 -e 100 -t 5</b>
Deletes messages with ids in range of 50 to 100 and 5 messages above.

