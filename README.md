# playfaster
A Gnome-Terminal Music Player written in bash that uses mpv. Playfaster has 3 processes running with the same terminal.

The first process (UI) is the user interface, it prints the different windows and is always waiting for user input (read).

The second process (Player) runs mpv one file at the time and WAITS for mpv to end. The playlists are created and managed by Playfaster and not by mpv. The Player knows the playlist contents and runs mpv. So it waits for mpv to end, and it's UI that stops mpv when the user changes the playing song. The two processes communicate via signals and "mails". Mails are short scripts containing data: when a process loads the script, data is passed between the two processes. 

The third process (Info) just asks mpv the time of the playing song and prints it to the screen. It does so many times each second.

Playfaster is very fast in printing the screen. It uses a single echo command every time it has to reprint a window. All tput commands, colors, texts, etc., are put in a single echo command.

The name Playfaster is because I use it for playing songs at a speed faster than normal, usually 1.2.

Playfaster uses awk for data manipolation. It has two data files: database.txt and playlists.txt. These are plain text files. awk is very fast in searching the files.

The main feature of Playfaster is that the user can create playlists with parameters for each song. The parameters are: Speed, Volume, Starting point, Ending point. Volume is a percentage value. So if you have an audio file where the volume is very low or very high compared to the other files, you can set this value. You can set a starting point for your audio file. It's very easy: when you listen to the song, you press the key to do that. Same thing for the ending point and the volume, and the speed.

Every time the user changes something, the UI writes a mail for the Player, then sends a signal (SIGUSR1), then, if necessary, stops mpv. The player checks if a mail has arrived with new parameters for playling the playlist.



License: Artistic License 2.0

https://opensource.org/license/Artistic-2.0
