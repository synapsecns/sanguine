let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/Code/Synapse/sanguine63
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +382 agents/agents/executor/executor.go
badd +321 agents/agents/guard/guard.go
badd +52 ~/Code/Synapse/sanguine63/agents/domains/domain.go
badd +130 ~/.config/nvim/init.lua
badd +356 ~/.tmux.conf.local
badd +689 agents/agents/notary/notary.go
badd +92 ~/Code/Synapse/sanguine63/agents/config/agent_config.go
badd +1 ethergo/submitter/config/config.go
badd +310 ethergo/submitter/suite_test.go
badd +142 agents/testutil/simulated_backends_suite.go
badd +32 ~/Code/Synapse/sanguine63/agents/agents/agentsintegration/suite_test.go
argglobal
%argdel
set stal=2
tabnew +setlocal\ bufhidden=wipe
tabrewind
edit ~/Code/Synapse/sanguine63/agents/agents/agentsintegration/suite_test.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 248 + 248) / 497)
exe 'vert 2resize ' . ((&columns * 248 + 248) / 497)
argglobal
balt agents/testutil/simulated_backends_suite.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 1 - ((0 * winheight(0) + 40) / 80)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 1
normal! 030|
wincmd w
argglobal
if bufexists(fnamemodify("ethergo/submitter/suite_test.go", ":p")) | buffer ethergo/submitter/suite_test.go | else | edit ethergo/submitter/suite_test.go | endif
if &buftype ==# 'terminal'
  silent file ethergo/submitter/suite_test.go
endif
balt agents/agents/guard/guard.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 119 - ((73 * winheight(0) + 40) / 80)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 119
normal! 011|
wincmd w
exe 'vert 1resize ' . ((&columns * 248 + 248) / 497)
exe 'vert 2resize ' . ((&columns * 248 + 248) / 497)
tabnext
edit ~/.config/nvim/init.lua
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 248 + 248) / 497)
exe 'vert 2resize ' . ((&columns * 248 + 248) / 497)
argglobal
balt agents/agents/guard/guard.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 130 - ((41 * winheight(0) + 40) / 80)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 130
normal! 0
wincmd w
argglobal
if bufexists(fnamemodify("~/.tmux.conf.local", ":p")) | buffer ~/.tmux.conf.local | else | edit ~/.tmux.conf.local | endif
if &buftype ==# 'terminal'
  silent file ~/.tmux.conf.local
endif
balt ~/.config/nvim/init.lua
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 209 - ((0 * winheight(0) + 40) / 80)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 209
normal! 024|
wincmd w
exe 'vert 1resize ' . ((&columns * 248 + 248) / 497)
exe 'vert 2resize ' . ((&columns * 248 + 248) / 497)
tabnext 1
set stal=1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let &winminheight = s:save_winminheight
let &winminwidth = s:save_winminwidth
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
