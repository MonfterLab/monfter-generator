## MONFTERS CLUB
This project is the tool for generate all 8000 Monfters.
To use this tool I assume that you have a basic golang knowledge.

#### The project structure
<pre>
├── assets(Original monfter trait)
│   ├── background
│   ├── body
│   ├── cap
│   ├── clothes
│   ├── eye
│   ├── horn
│   ├── mouth
│   └── nose
├── conf
├── dist
│   ├── images(The final monfters we see)
│   └── metadata(Monfters traits in json format)
├── scripts(Include sql and gen script)
</pre>

#### Generate Monfters on your computer
 - create two mysql tables names ```monfter_key``` and ```monfter_traits```
 - build the tool ```go build -o monfter-generator```
 - generate your own monfters ```./monfter-generator -c conf/app.conf -n 10```,   this may take a long time according to the param ```n``` which present the count you want to generate