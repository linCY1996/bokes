window.onload = function() {
    var np = new Vue({
        el: '#tall',
        data: {

        },
        methods: {

        },
        moounted: function() {

        }
    });
    $(document).ready(function() {

        //banner
        //  $('#banner').easyFader();
        $('#banner').easyFader();


        /*
      
        search
      
        */
        $('.search_ico').click(function() {
            $('.search_bar').toggleClass('search_open');
            if ($('#keyboard').val().length > 2) {
                $('#keyboard').val('');
                $('#searchform').submit();
            } else {
                return false;
            }
        });

        $('.cd-mao').click(function() {
            $("html, body").animate({
                scrollTop: $(".blogs3").offset().top + "px"
            }, {
                duration: 1000,
                easing: "swing"
            });
            return false;
        });
        $('.cd-mao1').click(function() {
            $("html, body").animate({
                scrollTop: $(".nav-margs-ch1").offset().top + "px"
            }, {
                duration: 1000,
                easing: "swing"
            });
            return false;
        });
        $('.cd-mao2').click(function() {
            $("html, body").animate({
                scrollTop: $(".nav-tall-ch2-right").offset().top + "px"
            }, {
                duration: 1000,
                easing: "swing"
            });
            return false;
        });
        $('.cd-mao3').click(function() {
            $("html, body").animate({
                scrollTop: $(".tell-me").offset().top + "px"
            }, {
                duration: 1000,
                easing: "swing"
            });
            return false;
        });
        $('.cd-mao4').click(function() {
            $("html, body").animate({
                scrollTop: $(".liuyan").offset().top + "px"
            }, {
                duration: 1000,
                easing: "swing"
            });
            return false;
        });
        //scroll to top
        var offset = 300,
            offset_opacity = 1200,
            scroll_top_duration = 700,
            $back_to_top = $('.cd-top');

        $(window).scroll(function() {
            ($(this).scrollTop() > offset) ? $back_to_top.addClass('cd-is-visible'): $back_to_top.removeClass('cd-is-visible cd-fade-out');
            if ($(this).scrollTop() > offset_opacity) {
                $back_to_top.addClass('cd-fade-out');
            }
        });
        $back_to_top.on('click', function(event) {
            event.preventDefault();
            $('body,html').animate({
                scrollTop: 0,
            }, scroll_top_duration);
        });

        //aside
        var Sticky = new hcSticky('aside', {
            stickTo: 'article',
            innerTop: 200,
            followScroll: false,
            queries: {
                480: {
                    disable: true,
                    stickTo: 'body'
                }
            }
        });

        //scroll
        if (!(/msie [6|7|8|9]/i.test(navigator.userAgent))) {
            window.scrollReveal = new scrollReveal({ reset: true });
        };



    });
}