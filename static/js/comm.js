window.onload = function() {
    var np = new Vue({
        el: '#tall',
        data: {
            firstmsg: '', //第二部分第一条数据
            msgsnext: [], // 第二部分剩下所有数据
            workmsgs: [], //程序员要问
            recommand: [], //特别推荐
            resnews: [], //最新推荐
            bowen: [], //博文
            timecar: [], //时间轴
            myself: '', //个人信息
            name: '',
            email: '',
            msgs: '',
            // usermsgs: [], //用户留言
        },
        methods: {
            showfir: function() {
                axios.get('/api/showfirstmsg', {
                    params: {
                        id: '1',
                    }
                }).then(function(resp) {
                    np.firstmsg = resp.data
                })
            },
            shownext: function() {
                axios.get('/api/showmsgnext').then(function(resp) {
                    np.msgsnext = resp.data
                })
            },
            showworkmsgs: function() {
                axios.get('/api/workmsgs').then(function(resp) {
                    np.workmsgs = resp.data
                })
            },
            showRecommand: function() {
                axios.get('/api/showrecommand').then(function(resp) {
                    np.recommand = resp.data
                })
            },
            showresnews: function() {
                axios.get('/api/showresnews').then(function(resp) {
                    np.resnews = resp.data
                })
            },
            showbowen: function() {
                axios.get('/api/showbowen').then(function(resp) {
                    np.bowen = resp.data
                })
            },
            showtimecar: function() {
                axios.get('/api/timecar').then(function(resp) {
                    np.timecar = resp.data
                })
            },
            showmyself: function() {
                axios.get('/api/myself').then(function(resp) {
                    np.myself = resp.data
                })
            },
            btn: function() {
                var that = this
                var params = new URLSearchParams();
                params.append('name', np.name);
                params.append('email', np.email);
                params.append('msgs', np.msgs);
                axios.post('/api/sendmsg', params).then(function(resp) {
                    location.reload();
                    alert(resp.data)
                })
            },
            // showusermsgs: function(resp) {
            //     axios.get('/api/usermsgs').then(function(resp) {
            //         np.usermsgs = resp.data
            //     })
            // },

        },
        mounted: function() {
            this.showfir();
            this.shownext();
            this.showworkmsgs();
            this.showRecommand();
            this.showresnews();
            this.showbowen();
            this.showtimecar();
            this.showmyself();
            // this.showusermsgs();
        }
    })
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