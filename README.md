# LPS
A simple program for counting the lines per second of stdin. This could be used to see the rate of growth of a log file or the rate of anything that prints to stdout. One common use case is piping tcpdump output to it to get the rate of filtered packets.

For example, to count the number of connections per second as identified by TCP SYN flags, you could use the following:

    tcpdump -i eth0 -n "tcp[tcpflags] & tcp-syn != 0" | ./lps


lps also accepts an interval parameter `-i` that allows you to specify the frequency at which the rate will be printed.
