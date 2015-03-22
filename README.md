## Download user agents 

Go to: [http://useragentstring.com/pages/All/](http://useragentstring.com/pages/All/)

Use: 

```
var jq = document.createElement('script');
jq.src = "https://ajax.googleapis.com/ajax/libs/jquery/1/jquery.min.js";
document.getElementsByTagName('head')[0].appendChild(jq);
// ... give time for script to load, then type.

$("#shaked").remove(); $("body").append("<div id='shaked'></div>"); $('ul li a').each(function(){ $t = $(this); var details = $t.parent().parent().prev().html().split(); var name = details[0]; var version = details[1]; $('#shaked').append("crc32.ChecksumIEEE([]byte(`" + $t.html() + "`)): Device{`" + name + "`,`" + version + "`},<br/>"); });


$("#shaked").remove(); 
$("body").append("<div id='shaked'></div>"); 
$('ul li a').each(function(){ 
    $t = $(this); 
    var details = $t.parent().parent().prev().html().split(); var name = details[0]; var version = details[1]; $('#shaked').append("crc32.ChecksumIEEE([]byte(`" + $t.html() + "`)): Device{`" + name + "`,`" + version + "`},<br/>"); });
```

Then copy the result that will show up in the textbox