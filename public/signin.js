
var memberDataInput = $('.member-data-input');
/*변수 선언*/
var todoListItem = $('.todo-list');

// var todoListInput = $('.id');
// var todoListInput2 = $('.todo-list-input-ppass');
// var todoListInput3 = $('.todo-list-input-bbirth');
// var todoListInput4 = $('.todo-list-input-ssex');


// var todoListInput5 = $('.todo-list-input-nname');
// var todoListInput6 = $('.todo-list-input-pphone');
// var todoListInput7 = $('.todo-list-input-eemail');






$('#btnJoin').on("click", function(event) {
    event.preventDefault();

    // var item = $(this).prevAll('.todo-list-input-iide').val();
    var item = $('#id').val();
    var pswd1 = $('#pswd1').val();
    var birthaa = $('#yy').val();
    var gender = $('#gender').val();
    var name = $('#name').val();
    var mobile = $('#mobile').val();
    var email = $('#email').val();
    // var mm = $('#mm').val();
    // var dd = $('#dd').val();
    // var birth = yyyy + '-' + mm + '-' + dd;
   
  
 
    // var item2 = $(this).prevAll('.todo-list-input-ppass').val();
    // var item3 = $(this).prevAll('.todo-list-input-bbirth').val();
    // var item4 = $(this).prevAll('.todo-list-input-ssex').val();

    // var item5 = $(this).prevAll('.todo-list-input-nname').val();
    // var item6 = $(this).prevAll('.todo-list-input-pphone').val();
    // var item7 = $(this).prevAll('.todo-list-input-eemail').val();
    


    if (item||item2||item3||item4||item5||item6||item7) {
        $.post("/todos", {ide:item,pass:pswd1,birth:birthaa,sex:gender,name:name,phone2:mobile,email:email}, addItem);
        // todoListInput.val("");
        // todoListInput2.val("");
        // todoListInput3.val("");
        
        // todoListInput4.val("");
        // todoListInput5.val("");
        // todoListInput6.val("");
        // todoListInput7.val("");

        memberDataInput.val("");
    }

});

var addItem = function(item) {
        todoListItem.append("<li "+ " id='" + item.id + "'><input class='checkbox' type='checkbox' />" + item.ide + "," + item.pass + "," + item.birth + "," + item.sex + "," + item.name + "," + item.phone2 + "," + item.email + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");           
};




$.get('/todos', function(items) {
    items.forEach(e => {
        addItem(e)
    
       
    });
});




todoListItem.on('change', '.checkbox', function() {
    var id = $(this).closest("li").attr('id');
    var $self = $(this);
    var complete = true;
    if ($(this).attr('checked')) {
        complete = false;
    }
    $.get("complete-todo/"+id+"?complete="+complete, function(data){
        if (complete) {
            $self.attr('checked', 'checked');
        } else {
            $self.removeAttr('checked');
        }

        $self.closest("li").toggleClass('completed');
    })
});

todoListItem.on('click', '.remove', function() {
    // url: todos/id method: DELETE
    var id = $(this).closest("li").attr('id');
    var $self = $(this);
    $.ajax({
        url: "todos/" + id,
        type: "DELETE",
        success: function(data) {
            if (data.success) {
                $self.parent().remove();
            }
        }
    })
    //$(this).parent().remove();
});

var id = document.querySelector('#id');

var pw1 = document.querySelector('#pswd1');
var pwMsg = document.querySelector('#alertTxt');
var pwImg1 = document.querySelector('#pswd1_img1');

var pw2 = document.querySelector('#pswd2');
var pwImg2 = document.querySelector('#pswd2_img1');
var pwMsgArea = document.querySelector('.int_pass');
// var pwMsgArea = document.querySelector('.todo-list-input-ppass');

var userName = document.querySelector('#name');

var yy = document.querySelector('#yy');
var mm = document.querySelector('#mm');
var dd = document.querySelector('#dd');

var gender = document.querySelector('#gender');

var email = document.querySelector('#email');

var mobile = document.querySelector('#mobile');

var error = document.querySelectorAll('.error_next_box');



/*이벤트 핸들러 연결*/


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
mobile.addEventListener("focusout", checkPhoneNum);





/*콜백 함수*/


function checkId() {
    var idPattern = /[a-zA-Z0-9_-]{5,20}/;
    if(id.value === "") {
        error[0].innerHTML = "필수 정보입니다.";
        error[0].style.display = "block";
    } else if(!idPattern.test(id.value)) {
        error[0].innerHTML = "5~20자의 영문 소문자, 숫자와 특수기호(_),(-)만 사용 가능합니다.";
        error[0].style.display = "block";
    } else {
        error[0].innerHTML = "멋진 아이디네요!";
        error[0].style.color = "#08A600";
        error[0].style.display = "block";
    }
}

function checkPw() {
    var pwPattern = /[a-zA-Z0-9~!@#$%^&*()_+|<>?:{}]{8,16}/;
    if(pw1.value === "") {
        error[1].innerHTML = "필수 정보입니다.";
        error[1].style.display = "block";
    } else if(!pwPattern.test(pw1.value)) {
        error[1].innerHTML = "8~16자 영문 대 소문자, 숫자, 특수문자를 사용하세요.";
        pwMsg.innerHTML = "사용불가";
        pwMsgArea.style.paddingRight = "93px";
        error[1].style.display = "block";
        
        pwMsg.style.display = "block";
        pwImg1.src = "m_icon_not_use.png";
    } else {
        error[1].style.display = "none";
        pwMsg.innerHTML = "안전";
        pwMsg.style.display = "block";
        pwMsg.style.color = "#03c75a";
        pwImg1.src = "m_icon_safe.png";
    }
}

function comparePw() {
    if(pw2.value === pw1.value && pw2.value != "") {
        pwImg2.src = "m_icon_check_enable.png";
        error[2].style.display = "none";
    } else if(pw2.value !== pw1.value) {
        pwImg2.src = "m_icon_check_disable.png";
        error[2].innerHTML = "비밀번호가 일치하지 않습니다.";
        error[2].style.display = "block";
    } 

    if(pw2.value === "") {
        error[2].innerHTML = "필수 정보입니다.";
        error[2].style.display = "block";
    }
}

function checkName() {
    var namePattern = /[a-zA-Z가-힣]/;
    if(userName.value === "") {
        error[3].innerHTML = "필수 정보입니다.";
        error[3].style.display = "block";
    } else if(!namePattern.test(userName.value) || userName.value.indexOf(" ") > -1) {
        error[3].innerHTML = "한글과 영문 대 소문자를 사용하세요. (특수기호, 공백 사용 불가)";
        error[3].style.display = "block";
    } else {
        error[3].style.display = "none";
    }
}


function isBirthCompleted() {
    var yearPattern = /[0-9]{4}/;

    if(!yearPattern.test(yy.value)) {
        error[4].innerHTML = "태어난 년도 4자리를 정확하게 입력하세요.";
        error[4].style.display = "block";
    } else {
        isMonthSelected();
    }


    function isMonthSelected() {
        if(mm.value === "월") {
            error[4].innerHTML = "태어난 월을 선택하세요.";
        } else {
            isDateCompleted();
        }
    }

    function isDateCompleted() {
        if(dd.value === "") {
            error[4].innerHTML = "태어난 일(날짜) 2자리를 정확하게 입력하세요.";
        } else {
            isBirthRight();
        }
    }
}



function isBirthRight() {
    var datePattern = /\d{1,2}/;
    if(!datePattern.test(dd.value) || Number(dd.value)<1 || Number(dd.value)>31) {
        error[4].innerHTML = "생년월일을 다시 확인해주세요.";
    } else {
        checkAge();
    }
}

function checkAge() {
    if(Number(yy.value) < 1920) {
        error[4].innerHTML = "정말이세요?";
        error[4].style.display = "block";
    } else if(Number(yy.value) > 2020) {
        error[4].innerHTML = "미래에서 오셨군요. ^^";
        error[4].style.display = "block";
    } else if(Number(yy.value) > 2005) {
        error[4].innerHTML = "만 14세 미만의 어린이는 보호자 동의가 필요합니다.";
        error[4].style.display = "block";
    } else {
        error[4].style.display = "none";
    }
}


function isEmailCorrect() {
    var emailPattern = /[a-z0-9]{2,}@[a-z0-9-]{2,}\.[a-z0-9]{2,}/;

    if(email.value === ""){ 
        error[6].style.display = "none"; 
    } else if(!emailPattern.test(email.value)) {
        error[6].style.display = "block";
    } else {
        error[6].style.display = "none"; 
    }

}

function checkPhoneNum() {
    var isPhoneNum = /([01]{2})([01679]{1})([0-9]{3,4})([0-9]{4})/;

    if(mobile.value === "") {
        error[7].innerHTML = "필수 정보입니다.";
        error[7].style.display = "block";
    } else if(!isPhoneNum.test(mobile.value)) {
        error[7].innerHTML = "형식에 맞지 않는 번호입니다.";
        error[7].style.display = "block";
    } else {
        error[7].style.display = "none";
    }

    
}
