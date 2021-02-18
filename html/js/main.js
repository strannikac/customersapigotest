const messageTime = 3000;
const effectsTime = 500;

const customersUrl = apiUrl + languageIso + '/customers/';

let screenWidth = 0;
let screenHeight = 0;

let isCustomersSearch = false;

function successFindCustomers(response) {
	if(response.html != undefined) {
		$('.js-customers .js-items tbody').html(response.html);
		$('.js-customers .js-paginator').html(response.paginator);
	}
}

function successSaveCustomer(response) {
	console.log(response);
    setTimeout(function() {
        $('.js-customer-form .js-frm-customer .error, .js-customer-form .js-frm-customer .success').hide();

		if(response.id != undefined && response.id > 0) {
			location.href = customersUrl + 'edit/' + response.id + '/';
		}
    }, messageTime);
}

function successDeleteCustomer(response) {
	location.reload();
}

function findCustomers(page, sort, desc, fname, lname) {
	const pages = +$('.js-customers .js-paginator a').last().attr('data-id');
	
	if(isNaN(page)) {
		page = 1;
	}
	if(sort == undefined) {
		sort = "";
	}
	if(desc == undefined) {
		desc = 0;
	}
	if(fname == undefined || fname.length < 2) {
		fname = "";
	}
	if(lname == undefined || lname.length < 2) {
		lname = "";
	}

	page = +page;
	if(page < 1) {
		page = 1;
	} else if(page > pages) {
		page = pages;
	}

	desc = +desc;
	if(desc != 1) {
		desc = 0;
	}

	let params = '';

	if(page > 1) {
		params = '?page=' + page;
	}
	if(sort != '' && sort != 'last_name') {
		params += params == '' ? '?' : '&';
		params += 'sort=' + sort;
	}
	if(desc == 1) {
		params += params == '' ? '?' : '&';
		params += 'desc=' + desc;
	}
	if(fname != '') {
		params += params == '' ? '?' : '&';
		params += 'fname=' + fname;
	}
	if(lname != '') {
		params += params == '' ? '?' : '&';
		params += 'lname=' + lname;
	}
	
	location.href = customersUrl + params;
}

function setPage() {
	screenHeight = +$(window).height();
	screenWidth = +$(window).width();
}

function setBirthdayDatepicker() {
    $(".js-datepicker.js-birthdate").datepicker({
        maxDate: 0
    });

	$('.js-datepicker.js-birthdate').datepicker("setDate", createDate($('.js-datepicker.js-birthdate').val()) );
}

function setCustmersParams() {
	let page = $.urlParam('page');
	let sort = $.urlParam('sort');
	let desc = $.urlParam('desc');
	let fname = $.urlParam('fname');
	let lname = $.urlParam('lname');

	if(page == null || page == 0) {
		page = 1;
	}

	if(sort == null || sort == 0) {
		sort = "last_name";
	}

	if(desc == null || desc != "1") {
		desc = 0;
	}

	if(fname == null || fname == 0) {
		fname = "";
	}

	if(lname == null || lname == 0) {
		lname = "";
	}

	let el = $('.js-customers .js-items th a[data-sort="' + sort + '"]');

	if(el.attr('data-sort') == undefined) {
		sort = 'last_name';
		el = $('.js-customers .js-items th a[data-sort="' + sort + '"]');
	}
	
	$('.js-customers .js-items th a').removeClass('active');
	$('.js-customers .js-items th a span').html('');
	el.addClass('active');

	$('.js-customers .js-items th a').attr('data-desc', 0);
	el.attr('data-desc', desc);

	const arrow = desc == 1 ? '&uarr;' : '&darr;';
	el.find('span').html(arrow);

	if(fname != '' || lname != '') {
		$('.js-customers .js-frm-search input[name=fname]').val(fname);
		$('.js-customers .js-frm-search input[name=lname]').val(lname);
		isCustomersSearch = true;
	}
}

$(function(){
	ARR_LOCALES = JSON.parse(ARR_LOCALES);

	$(window).resize(setPage);
	setPage();

	if($('.js-customers .js-items').attr('class') != undefined) {
		setCustmersParams();
	}

	if($('.js-customer-form .js-frm-customer').attr('class') != undefined) {
		setBirthdayDatepicker();
	}

    $('body').on("click", '.js-lang-switcher .js-current', function(){
        $('.js-lang-switcher ul').css('display', 'flex');
		return false;
    });

    $('body').on("click", '.js-lang-switcher ul li a', function(){
		const iso = $(this).attr('data-iso');

		if(iso != languageIso) {
			let link = location.href;
			if(link.indexOf('/' + languageIso + '/') > 0) {
				link = link.replace('/' + languageIso + '/', '/' + iso + '/');
			} else {
				link = $(this).attr('href');
			}
			location.href = link;
		}

        return false;
    });
    
    $('body').on("click", '.js-customers .js-items th a', function(){
        let desc = 0;
        if($(this).hasClass('active')) {
            if($(this).attr('data-desc') != 1) {
                desc = 1;
            }
        }

		let page = +$('.js-customers .js-paginator a.active').first().attr('data-id');

		let fname = "";
		let lname = "";

		if(isCustomersSearch) {
			fname = $('.js-frm-search input[name=fname]').first().val();
			lname = $('.js-frm-search input[name=lname]').first().val();
		}

        findCustomers(page, $(this).attr('data-sort'), desc, fname, lname);

        return false;
    });
    
    $('body').on("click", '.js-customers .js-paginator a', function(){
		if(!$(this).hasClass('active')) {
			const sort = $('.js-customers .js-items th a.active').attr('data-sort');
			const desc = $('.js-customers .js-items th a.active').attr('data-desc');
			let fname = "";
			let lname = "";

			if(isCustomersSearch) {
				fname = $('.js-frm-search input[name=fname]').first().val();
				lname = $('.js-frm-search input[name=lname]').first().val();
			}

			findCustomers($(this).attr('data-id'), sort, desc, fname, lname);
		}

        return false;
    });

	$('body').on("submit", '.js-customers .js-frm-search', function(){
		let fname = $(this).find('input[name=fname]').val();
		let lname = $(this).find('input[name=lname]').val();
		let error = '';

		$(this).find('input').removeClass('err-el');
		$(this).find('.error').hide();

		if(fname == '' && lname == '') {
			error = ARR_LOCALES['ERR_EMPTY_SEARCH_DATA'];
			$(this).find('input').addClass('err-el');
		} else if(fname != '' && fname.length < 3) {
			error = ARR_LOCALES['ERR_MIN_SEARCH_LENGTH'];
			$(this).find('input[name=fname]').addClass('err-el');
		} else if(lname != '' && lname.length < 3) {
			error = ARR_LOCALES['ERR_MIN_SEARCH_LENGTH'];
			$(this).find('input[name=lname]').addClass('err-el');
		}

		if(error != '') {
			$(this).find('.error').html(error).show();
			return false;
		}

		isCustomersSearch = true;

		const sort = $('.js-customers .js-items th a.active').attr('data-sort');
		const desc = $('.js-customers .js-items th a.active').attr('data-desc');

		findCustomers(1, sort, desc, fname, lname);

        return false;
    });
    
    $('body').on("click", '.js-customers .js-link-search-clear', function(){
		$('.js-customers .js-link-search-clear').parents('.js-frm-search').find('input').val('');
		isCustomersSearch = false;

		findCustomers();

        return false;
    });

	$('body').on('submit', '.js-customer-form .js-frm-customer', function(){
		const id = +$(this).find('input[name=id]').val();

		let method = 'POST';
		let url = customersUrl;
		if(id > 0) {
			method = 'PUT';
			url += id + '/';
		}

		sendForm( '.js-customer-form .js-frm-customer', method, url, 'successSaveCustomer' );

        return false;
    });
    
    $('body').on("click", '.js-customers .js-items  .js-link-delete', function(){
		let id = +$(this).parents('tr').attr('data-id');

		sendAjax( '', 'DELETE', customersUrl + id + '/', '', 'successDeleteCustomer' );

        return false;
    });
});