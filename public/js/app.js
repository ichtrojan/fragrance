
$('.sp-grid').slick({
    infinite: false,
    slidesToShow: 4,
    slidesToScroll: 1,
    responsive: [
        {
            breakpoint: 1024,
            settings: {
                slidesToShow: 3,

            }
        },
        {
            breakpoint: 768,
            settings: "unslick"
        },

        // You can unslick at a given breakpoint now by adding:
        // settings: "unslick"
        // instead of a settings object
    ]
})

$('.pp-grid').slick({
    infinite: false,
    slidesToShow: 3,
    slidesToScroll: 1,
    responsive: [
        {
            breakpoint: 1024,
            settings: {
                slidesToShow: 2,

            }
        },
        {
            breakpoint: 768,
            settings: "unslick"
        },
        // You can unslick at a given breakpoint now by adding:
        // settings: "unslick"
        // instead of a settings object
    ]
})

$('#minus').click(function (e) {
    e.preventDefault()
    let value = $('#amount').val()
    let price = $('input:radio.sizes:checked').data('price')
    let single_price = $('#total-price').data('single')


    if (value != 0) {
        value--
        $('#amount').val(value)
        var total = (parseFloat(single_price) + parseFloat(price))
        $('#total-price').data('price', total)
        $('#total-price').text(total * value)

    }
})


$('#plus').click(function (e) {
    e.preventDefault()
    if ($('input:radio.sizes:checked').val() == null) {
        alert('select a size first')
        return
    }
    let value = $('#amount').val()
    let price = $('input:radio.sizes:checked').data('price')
    let single_price = $('#total-price').data('single')


    value++
    $('#amount').val(value)
    var total = (parseFloat(single_price) + parseFloat(price))
    $('#total-price').data('price', total)
    $('#total-price').text(total * value)

})

$('.back').click(function () {
    window.history.back()
})

$('input:radio.sizes').change(function (e) {
        let price = $(this).data('price')
        let value = $('#amount').val()
        let single_price = $('#total-price').data('single')

        var total = (parseFloat(single_price) + parseFloat(price))
        $('#total-price').data('price', total)
        $('#total-price').text(total * value)

})

$('#perfume').imgcolr(function(img, color){
    $('.cp-sidebar').css('background', `linear-gradient(270deg,${color} 70%,#ffffff 87.5%)`)
})