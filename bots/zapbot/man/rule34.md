<b>rule34</b> [<u>options</u>] [<u>tags</u>]

<code>-c</code> <u>number</u> 
Number of posts you want to have. From 2 to 10

<code>-r</code>
Selectes random post from first 1000 ones

<code>-p</code> <u>number</u>
Page number. From 1

<b>EXAMPLES</b>

<b>rule34 -c 10</b>
Shows the newest 10 posts

<b>rule34 -r -c 10</b>
Picks randomly 10 posts from the newest 1000 posts

<b>rule34 -p 1 -c 5</b>
Skips first 100 posts, and shows first 5 posts from the 2nd hundred

<b>rule34 -p 3 -c 8</b>
Skips first 300 posts, and shows first 8 posts from the 4th hundred

<b>rule34 -r -p 5</b>
Skips first 5000 posts, and picks randomly 8 posts from the 6th thousand

<b>rule34</b> <u>tags</u> 
Shows the newest post with given tags. Tags has a format of snake case, e.x. hug_from_behind, female_only. If you want to get a specific character from a specific fandom use the next format: character_name_(fandom_name), e.x. hu_tao_(genshin_impact). Also you can write multiple of them, separating by space, and posts will definitely have every of the given tag

<b>rule34 -r -c 10 sfw wholesome 2girls</b>
Will randomly pick 10 posts from the first 1000 posts and every post will contain next tags: sfw wholesome 2girls

