Get server links (https://vipstream.tv/ajax/v2/tv/seasons/ <- filmid)
Via parsing this html (example)

```html
<body><div class="seasons-list seasons-list-new border-bottom-block">
    <div id="servers-list"></div>
    <div class="sl-content">
        <div class="slc-eps">
            <div class="sl-title">
                <div class="slt-seasons-dropdown">
                    <button class="btn btn-seasons dropdown-toggle" type="button" data-toggle="dropdown"><i class="fas fa-list mr-2"></i><span id="current-season"></span>
                    </button>
                    <div class="dropdown-menu dropdown-menu-new">
                        
                            <a data-id="385" id="ss-385" data-toggle="tab" class="dropdown-item ss-item" href="#ss-episodes-385">Season 1</a>
                        
                            <a data-id="386" id="ss-386" data-toggle="tab" class="dropdown-item ss-item" href="#ss-episodes-386">Season 2</a>
                        
                            <a data-id="387" id="ss-387" data-toggle="tab" class="dropdown-item ss-item" href="#ss-episodes-387">Season 3</a>
                        
                            <a data-id="388" id="ss-388" data-toggle="tab" class="dropdown-item ss-item" href="#ss-episodes-388">Season 4</a>
                        
                            <a data-id="389" id="ss-389" data-toggle="tab" class="dropdown-item ss-item" href="#ss-episodes-389">Season 5</a>
                        
                    </div>
                </div>
            </div>
            <div class="clearfix"></div>
            <div class="slce-list">
                <div class="tab-content">
                    
                        <div class="ss-episodes tab-pane fade" id="ss-episodes-385"></div>
                    
                        <div class="ss-episodes tab-pane fade" id="ss-episodes-386"></div>
                    
                        <div class="ss-episodes tab-pane fade" id="ss-episodes-387"></div>
                    
                        <div class="ss-episodes tab-pane fade" id="ss-episodes-388"></div>
                    
                        <div class="ss-episodes tab-pane fade" id="ss-episodes-389"></div>
                    
                </div>
            </div>
            <div class="clearfix"></div>
        </div>
        <div class="clearfix"></div>
    </div>
</div>
<script>
    $('.ss-item').click(function () {
        $("#servers-list").empty();
        $('.ss-item').removeClass('active');
        $('.ss-episodes').removeClass('active show');
        $(this).addClass('active');
        $('#current-season').text($(this).text());

        var ssID = $(this).attr('data-id');
        $('#ss-episodes-' + ssID).addClass('active show');

        if ($('#ss-episodes-' + ssID).is(':empty')) {
            $.get("/ajax/v2/season/episodes/" + ssID, function (res) {
                $('#ss-episodes-' + ssID).html(res);
                chooseEpisode(ssID);
            });
        } else {
            chooseEpisode(ssID);
        }
    });

    function chooseEpisode(ssID) {
        var currentEpsID = $('.detail_page-watch').attr('data-episode');
        if ($('#episode-' + currentEpsID).length > 0) {
            $('#episode-' + currentEpsID).click();
            $('.detail_page-watch').attr('data-episode', '');
        } else {
            if ($('.detail_page').hasClass('watch_page')) {
                $('#ss-episodes-' + ssID + ' .eps-item').first().click();
            }
        }
    }

    var ssID = $('.detail_page-watch').attr('data-season');
    if (ssID) {
        $('#ss-' + ssID).click();
        $('#ss-episodes-' + ssID).addClass('active show');
    } else {
        $('.ss-item').first().click();
        $('.ss-episodes').first().addClass('active show');
    }
</script></body></html>
```
parse the html for the data-id

via associating with season

season 1 -> 385
season 2 -> 386

https://vipstream.tv/ajax/v2/season/episodes/ <- 385

Getting episodes

```html
<body><ul class="nav">
    
        <li class="nav-item">
            <a id="episode-7013" data-id="7013" class="nav-link btn btn-sm btn-secondary eps-item" href="javascript:;" title="Eps 1: Pilot"><i class="fas fa-play"></i><strong>Eps 1
                    :</strong> Pilot</a>
        </li>
    
        <li class="nav-item">
            <a id="episode-7014" data-id="7014" class="nav-link btn btn-sm btn-secondary eps-item" href="javascript:;" title="Eps 2: The Cat's in the Bag"><i class="fas fa-play"></i><strong>Eps 2
                    :</strong> The Cat's in the Bag</a>
        </li>
    
        <li class="nav-item">
            <a id="episode-7015" data-id="7015" class="nav-link btn btn-sm btn-secondary eps-item" href="javascript:;" title="Eps 3: And the Bag's in the River"><i class="fas fa-play"></i><strong>Eps 3
                    :</strong> And the Bag's in the River</a>
        </li>
    
        <li class="nav-item">
            <a id="episode-7016" data-id="7016" class="nav-link btn btn-sm btn-secondary eps-item" href="javascript:;" title="Eps 4: Cancer Man"><i class="fas fa-play"></i><strong>Eps 4
                    :</strong> Cancer Man</a>
        </li>
    
        <li class="nav-item">
            <a id="episode-7017" data-id="7017" class="nav-link btn btn-sm btn-secondary eps-item" href="javascript:;" title="Eps 5: Gray Matter"><i class="fas fa-play"></i><strong>Eps 5
                    :</strong> Gray Matter</a>
        </li>
    
        <li class="nav-item">
            <a id="episode-7018" data-id="7018" class="nav-link btn btn-sm btn-secondary eps-item" href="javascript:;" title="Eps 6: Crazy Handful of Nothin'"><i class="fas fa-play"></i><strong>Eps 6
                    :</strong> Crazy Handful of Nothin'</a>
        </li>
    
        <li class="nav-item">
            <a id="episode-7019" data-id="7019" class="nav-link btn btn-sm btn-secondary eps-item" href="javascript:;" title="Eps 7: A No-Rough-Stuff-Type Deal"><i class="fas fa-play"></i><strong>Eps 7
                    :</strong> A No-Rough-Stuff-Type Deal</a>
        </li>
    
</ul>
```

https://vipstream.tv/ajax/v2/episode/servers/7019/#servers-list

parse html for server id

```html
<html><head></head><body><div class="detail_page-servers mb-0">
    <div class="dp-s-line">
        <div class="server-notice text-center">
            <span>If current server doesn't work please try other servers below.</span></div>
        <ul class="nav">
            
                <li class="nav-item">
                    <a data-id="4858930" id="watch-4858930" class="nav-link btn btn-sm btn-secondary link-item active" href="javascript:;" title="Server UpCloud"><i class="fas fa-play mr-2"></i><span>UpCloud</span></a>
                </li>
            
                <li class="nav-item">
                    <a data-id="1615437" id="watch-1615437" class="nav-link btn btn-sm btn-secondary link-item " href="javascript:;" title="Server Vidcloud"><i class="fas fa-play mr-2"></i><span>Vidcloud</span></a>
                </li>
            
                <li class="nav-item">
                    <a data-id="6140050" id="watch-6140050" class="nav-link btn btn-sm btn-secondary link-item " href="javascript:;" title="Server Streamlare"><i class="fas fa-play mr-2"></i><span>Streamlare</span></a>
                </li>
            
                <li class="nav-item">
                    <a data-id="1064154" id="watch-1064154" class="nav-link btn btn-sm btn-secondary link-item " href="javascript:;" title="Server Upstream"><i class="fas fa-play mr-2"></i><span>Upstream</span></a>
                </li>
            
                <li class="nav-item">
                    <a data-id="5684935" id="watch-5684935" class="nav-link btn btn-sm btn-secondary link-item " href="javascript:;" title="Server MixDrop"><i class="fas fa-play mr-2"></i><span>MixDrop</span></a>
                </li>
            
        </ul>
    </div>
</div>
```

query the ajax one last time to get the provider link after that get the sources for that provider

https://vipstream.tv/ajax/get_link/5684935
