
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
    if (value != 0) {
        value--
        $('#amount').val(value)
    }
})


$('#plus').click(function (e) {
    e.preventDefault()
    let value = $('#amount').val()
    value++
    $('#amount').val(value)
})