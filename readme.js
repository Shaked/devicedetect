$("#shaked").remove();
$("body").append("<textarea id='shaked'></textarea>");
$('ul li a').each(function() {
    $t = $(this);
    var details = $t.parent().parent().prev().html().split();
    var name = details[0];
    var version = details[1];
    $('#shaked').append("crc32.ChecksumIEEE([]byte(`" + $t.html() + "`)): Device{`" + name + "`,`" + version + "`},<br/>");
});


var jq = document.createElement('script');
jq.src = "https://ajax.googleapis.com/ajax/libs/jquery/1/jquery.min.js";
document.getElementsByTagName('head')[0].appendChild(jq);
if ($("#shaked").length > 0) {
    $("#shaked").remove();
}
$("body").append("<textarea id='shaked'></textarea>");
var text = '';
$('ul li a').each(function() {
    $t = $(this);
    var details = $t.parent().parent().prev().html().split();
    var name = details[0];
    var version = details[1];
    text += "`" + $t.html() + "`: `" + name + "`,\n"
});
$("#shaked").text(text)