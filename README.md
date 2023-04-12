
# ascii-art

## Description

Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. In this project use banner standard.txt

Banner Format

- Each character has a height of 8 lines.
- Characters are separated by a new line `\n`.
- Here is an example of ' ', '!' and '"'(one dot represents one space):

```console
......
......
......
......
......
......
......
......

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....

._._..
(.|.).
.V.V..
......
......
......
......
......
```

## Usage

```console
student$ go run . "" | cat -e
student$ go run . "\n" | cat -e
$
student$ go run . "Hello\n" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
student$ go run . "HeLlO" | cat -e
 _    _          _        _    ____   $
| |  | |        | |      | |  / __ \  $
| |__| |   ___  | |      | | | |  | | $
|  __  |  / _ \ | |      | | | |  | | $
| |  | | |  __/ | |____  | | | |__| | $
|_|  |_|  \___| |______| |_|  \____/  $
                                      $
                                      $

student$ go run . "1Hello 2There" | cat -e
     _    _          _   _                      _______   _                           $
 _  | |  | |        | | | |              ____  |__   __| | |                          $
/ | | |__| |   ___  | | | |   ___       |___ \    | |    | |__     ___   _ __    ___  $
| | |  __  |  / _ \ | | | |  / _ \        __) |   | |    |  _ \   / _ \ | '__|  / _ \ $
| | | |  | | |  __/ | | | | | (_) |      / __/    | |    | | | | |  __/ | |    |  __/ $
|_| |_|  |_|  \___| |_| |_|  \___/      |_____|   |_|    |_| |_|  \___| |_|     \___| $
                                                                                      $
                                                                                      $
student$ go run . "{Hello There}" | cat -e
   __  _    _          _   _               _______   _                           __    $
  / / | |  | |        | | | |             |__   __| | |                          \ \   $
 | |  | |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___   | |  $
/ /   |  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \   \ \ $
\ \   | |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/   / / $
 | |  |_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___|  | |  $
  \_\                                                                            /_/   $
                                                                                       $
student$ go run . "Hello\nThere" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
student$ go run . "Hello\n\nThere" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
```

## Implementation details

- read file to []byte
- split string into substrings separated by \n
- append into array if len of string not equal 0
- if len of string equal 0, append array into map with key (first symbol = " "). Then reset array, first symbol increase by one
- then print char, find symbol in key of map then print first line, second line, ... eighth line.

