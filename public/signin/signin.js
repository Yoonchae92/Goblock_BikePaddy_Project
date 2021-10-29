(function($) {
    'use strict';
    $(function() {
        var memberItem = $('.member-data-list');
        var memberDataInput = $('.member-data-input');
        

    /*변수 선언*/


    var id = document.querySelector('#id');

    var pw1 = document.querySelector('#pswd1');
    var pwMsg = document.querySelector('#alertTxt');
    var pwImg1 = document.querySelector('#pswd1_img1');

    var pw2 = document.querySelector('#pswd2');
    var pwImg2 = document.querySelector('#pswd2_img1');
    var pwMsgArea = document.querySelector('.int_pass');

    var userName = document.querySelector('#name');

    var yy = document.querySelector('#yy');
    var mm = document.querySelector('#mm');
    var dd = document.querySelector('#dd');

    var gender = document.querySelector('#gender');

    var email = document.querySelector('#email');

   // var mobile = document.querySelector('#mobile');

    var error = document.querySelectorAll('.error_next_box');





    

    /*이벤트 핸들러 연결*/
    area.addEventListener("focusout", checkarea);
     bike_info.addEventListener("focusout", checkbike_info);
    // career.addEventListener("focusout", checkcareer);
     club.addEventListener("focusout", checkclub);

    id.addEventListener("focusout", checkId);
    pw1.addEventListener("focusout", checkPw);
    pw2.addEventListener("focusout", comparePw);
    userName.addEventListener("focusout", checkName);
    yy.addEventListener("focusout", isBirthCompleted);
    mm.addEventListener("focusout", isBirthCompleted);
    dd.addEventListener("focusout", isBirthCompleted);
    gender.addEventListener("focusout", function() {
        if(gender.value === "성별") {
            error[5].style.display = "block";
        } else {
            error[5].style.display = "none";
        }
    })
    email.addEventListener("focusout", isEmailCorrect);
    //mobile.addEventListener("focusout", checkPhoneNum);


    function checkarea() {
    }
    function checkbike_info() {
    }
    function checkcareer() {
    }
    function checkclub() {
    }

    /*콜백 함수*/


    function checkId() {
        var idPattern = /[a-zA-Z0-9_-]{5,20}/;
        if(id.value === "") {
            error[0].innerHTML = "필수 정보입니다.";
            error[0].style.display = "block";
            return false;
        } else if(!idPattern.test(id.value)) {
            error[0].innerHTML = "5~20자의 영문 소문자, 숫자와 특수기호(_),(-)만 사용 가능합니다.";
            error[0].style.display = "block";
            return false;
        } else {
            error[0].innerHTML = "멋진 아이디네요!";
            error[0].style.color = "#08A600";
            error[0].style.display = "block";
            return true;
        }
    }

    function checkPw() {
        var pwPattern = /[a-zA-Z0-9~!@#$%^&*()_+|<>?:{}]{8,16}/;
        if(pw1.value === "") {
            error[1].innerHTML = "필수 정보입니다.";
            error[1].style.display = "block";
            return false;
        } else if(!pwPattern.test(pw1.value)) {
            error[1].innerHTML = "8~16자 영문 대 소문자, 숫자, 특수문자를 사용하세요.";
            pwMsg.innerHTML = "사용불가";
            pwMsgArea.style.paddingRight = "93px";
            error[1].style.display = "block";
            
            pwMsg.style.display = "block";
            pwImg1.src = "m_icon_not_use.png";
            return false;
        } else {
            error[1].style.display = "none";
            pwMsg.innerHTML = "안전";
            pwMsg.style.display = "block";
            pwMsg.style.color = "#03c75a";
            pwImg1.src = "m_icon_safe.png";
            return true;
        }
    }

    function comparePw() {
        if(pw2.value === pw1.value && pw2.value != "") {
            pwImg2.src = "m_icon_check_enable.png";
            error[2].style.display = "none";
            return true;
        } else if(pw2.value !== pw1.value) {
            pwImg2.src = "m_icon_check_disable.png";
            error[2].innerHTML = "비밀번호가 일치하지 않습니다.";
            error[2].style.display = "block";
            return false;
        } 

        if(pw2.value === "") {
            error[2].innerHTML = "필수 정보입니다.";
            error[2].style.display = "block";
            return false;
        }
    }

    function checkName() {
        var namePattern = /[a-zA-Z가-힣]/;
        if(userName.value === "") {
            error[3].innerHTML = "필수 정보입니다.";
            error[3].style.display = "block";
            return false;
        } else if(!namePattern.test(userName.value) || userName.value.indexOf(" ") > -1) {
            error[3].innerHTML = "한글과 영문 대 소문자를 사용하세요. (특수기호, 공백 사용 불가)";
            error[3].style.display = "block";
            return false;
        } else {
            error[3].style.display = "none";
            return true;
        }
    }


    function isBirthCompleted() {
        var yearPattern = /[0-9]{4}/;

        if(!yearPattern.test(yy.value)) {
            error[4].innerHTML = "태어난 년도 4자리를 정확하게 입력하세요.";
            error[4].style.display = "block";
            return false;
        } else {
            return isMonthSelected();
        }


        function isMonthSelected() {
            if(mm.value === "월") {
                error[4].innerHTML = "태어난 월을 선택하세요.";
                return false;
            } else {
                return isDateCompleted();
            }
        }

        function isDateCompleted() {
            if(dd.value === "") {
                error[4].innerHTML = "태어난 일(날짜) 2자리를 정확하게 입력하세요.";
                return false;
            } else {
                return isBirthRight();
            }
        }
    }



    function isBirthRight() {
        var datePattern = /\d{1,2}/;
        if(!datePattern.test(dd.value) || Number(dd.value)<1 || Number(dd.value)>31) {
            error[4].innerHTML = "생년월일을 다시 확인해주세요.";
            return false;
        } else {
            return checkAge();
        }
    }

    function checkAge() {
        if(Number(yy.value) < 1920) {
            error[4].innerHTML = "정말이세요?";
            error[4].style.display = "block";
            return false;
        } else if(Number(yy.value) > 2020) {
            error[4].innerHTML = "미래에서 오셨군요. ^^";
            error[4].style.display = "block";
            return false;
        } else if(Number(yy.value) > 2005) {
            error[4].innerHTML = "만 14세 미만의 어린이는 보호자 동의가 필요합니다.";
            error[4].style.display = "block";
            return false;
        } else {
            error[4].style.display = "none";
            return true;
        }
    }


    function isEmailCorrect() {
        var emailPattern = /[a-z0-9]{2,}@[a-z0-9-]{2,}\.[a-z0-9]{2,}/;

        if(email.value === ""){ 
            error[6].style.display = "none";
            return true; 
        } else if(!emailPattern.test(email.value)) {
            error[6].style.display = "block";
            return false;
        } else {
            error[6].style.display = "none"; 
            return true;
        }

    }

    // function checkPhoneNum() {
    //     var isPhoneNum = /([01]{2})([01679]{1})([0-9]{3,4})([0-9]{4})/;

    //     if(mobile.value === "") {
    //         error[7].innerHTML = "필수 정보입니다.";
    //         error[7].style.display = "block";
    //         return false;
    //     } else if(!isPhoneNum.test(mobile.value)) {
    //         error[7].innerHTML = "형식에 맞지 않는 번호입니다.";
    //         error[7].style.display = "block";
    //         return false;
    //     } else {
    //         error[7].style.display = "none";
    //         return true;
    //     }

        
    // }


    /*
    2월 : 윤년에는 29일까지, 평년에는 28일까지.
    1,3,5,7, 8,10,12 -> 31일
    2,4,6, 9,11 -> 30일

        var days31 = [1, 3, 5, 7, 8, 10, 12];
        var days30 = [4, 6, 9, 11];

        if(mm.value )

    var sel = document.getElementById("sel");
    var val = sel.options[sel.selectedIndex].value;

    var id = document.querySelector('#id');
    var pw1 = document.querySelector('#pswd1');
    var pw2 = document.querySelector('#pswd2');
    var yourName = document.querySelector('#name');
    var yy = document.querySelector('#yy');
    var mm = document.querySelector('#mm');
    var dd = document.querySelector('#dd');
    var email = document.querySelector('#email');
    var mobile = document.querySelector('#mobile');
    var error = document.querySelectorAll('.error_next_box');

    var pattern_num = /[0-9]/;
    var pattern_spc = /[~!@#$%^&*()_+|<>?:{}]/;


    id.onchange = checkId;
    pw1.onchange = checkPw;
    pw2.onchange = comparePw;
    yourName.onchange = checkName;
    yy.onchange = checkYear;


    function checkId() {
        if(id.value === "") {
            error[0].style.display = "block";
        } else if(id.value.length < 5 || id.value.length > 20){
            error[0].innerHTML = "5~20자의 영문 소문자, 숫자와 특수기호(_),(-)만 사용 가능합니다.";
            error[0].style.display = "block";
        }
    }

    function checkPw() {
        if(pw1.value === "") {
            error[1].style.display = "block";
        } else if (pw1.value.length < 8 || pw1.value.length > 16) {
            error[1].innerHTML = "8~16자 영문 대 소문자, 숫자, 특수문자를 사용하세요.";
            error[1].style.display = "block";
        }
    }

    function comparePw() {
        if(pw2.value === "") {
            error[2].style.display = "block";
        } else if (pw2.value !== pw1.value) {
            error[2].innerHTML = "비밀번호가 일치하지 않습니다.";
            error[2].style.display = "block";
        }
    }

    function checkName() {
        if( yourName.value.indexOf(" ") >= 0 || pattern_spc.test(yourName.value) || pattern_num.test(yourName.value) ) {
            error[3].innerHTML = "한글과 영문 대 소문자를 사용하세요. (특수기호, 공백 사용 불가)";
            error[3].style.display = "block";
        } else if(yourName.value.length === 0) {
            error[3].style.display = "block";
        } else {
            error[3].style.display = "none";
        }
    }

    function checkYear() {
        isBirthEntered();
        if(yy.value.length !== 4 || !pattern_num.test(yy.value)) {
            error[4].innerHTML = "태어난 년도 4자리를 정확하게 입력하세요.";
        } else if (parseInt(yy.value) < 1920) {
            error[4].innerHTML = "정말이세요?";
            error[4].style.display = "block";
        }

    }

    function isBirthEntered() {
        
    }

    function checkEmail() {
        
    }

    function checkNumber() {
        
    }
    */
        
        $('#btnJoin').on("click", function(event) {
            event.preventDefault();

            var id = $('#id').val();
            var pswd1 = $('#pswd1').val();
            var name = $('#name').val();
            var yyyy = $('#yy').val();
            var mm = $('#mm').val();
            var dd = $('#dd').val();
            var birth = yyyy + '-' + mm + '-' + dd;
            var gender = $('#gender').val();
            var email = $('#email').val();
          //  var mobile = $('#mobile').val();
            var club = $('#club').val();

            var area = $('#area').val();
            var bike_info = $('#bike_info').val();
            var career = $('#career').val();

            if (checkId() || checkPw() || comparePw() || checkName() || isBirthCompleted() || isEmailCorrect() || checkclub()|| checkarea()|| checkbike_info()|| checkcareer()) {
                $.post("/members", {id : id, pswd : pswd1, name : name, gender : gender, birth : birth, email : email, club : club, area :area,bike_info:bike_info,career:career}, addItem)
                //memberItem.append("<li><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
                memberDataInput.val("");
            }
        });



        var addItem = function(item) {
            if (item.completed) {
                memberItem.append("<li class='completed'" + " id='" + item.id + "'>" + item.id + "," + item.pswd + "," + item.name + "," + item.birth + "," + item.gender + "," + item.email + "," + item.club + "," + item.area + "," + item.bike_info + "," + item.career + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
            } else {
                memberItem.append("<li " + " id='" + item.id + "'>" + item.id + "," + item.pswd + "," + item.name + "," + item.birth + "," + item.gender + "," + item.email + "," + item.club + "," + item.area + "," + item.bike_info + "," + item.career + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
            }
        };
// 2단계
        // $.get('/members', function(items) {
        //     items.forEach(e => {
        //         addItem(e)
        //     });
        // });


        memberItem.on('click', '.remove', function() {
            // url: todos/id method: DELETE
            var id = $(this).closest("li").attr('id');
            var $self = $(this);
            $.ajax({
                    url: "members/" + id,
                    type: "DELETE", 
                    success: function(data) {
                        if(data.success) {
                            $self.parent().remove();
                        }
                    }
                })
            //$(this).parent().remove();
        });

    });
})(jQuery);