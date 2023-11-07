// adding dashes to the phone numbers
const phoneNumberSpanElems = document.querySelectorAll("span.phone-number-text")
phoneNumberSpanElems.forEach((baseElem) => {
    baseElem.innerText = "+1-" + baseElem.innerText.replace(/(\d{3})(\d{3})(\d{4})/, "$1-$2-$3")
})