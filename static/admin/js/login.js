function getCaptcha() {
    // 加上时间戳，防止浏览器缓存
    fetch("/admin/getCaptcha?t="+Date.now()).then(res => res.json()).then(data => {
        document.getElementById("captchaId").value = data.captchaId
        document.getElementById("captchaImg").src = data.captchaImage
    })
}

getCaptcha()