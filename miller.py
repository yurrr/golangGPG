#!/usr/bin/python2
 
nqtde = input()
i = 0
 
 
def potMod(mod, base, exp, chamada):
    j = 1 
 
    while exp > 0 :
        if exp%2 == 1 :
           # if chamada == 1:
           #     print j, base, exp, "S"
            exp = (exp-1)/2
            j = (j*base) % mod
        else :
           # if chamada == 1:
           #     print j, base, exp, "N"
            exp = exp/2
        base = (base*base) %mod
 
   
        
       # if chamada == 1:
       # print j, base, exp, "N" 
 
    else: 
        return j
 
 
 
def header(num):
     teste = num -1  
     ctrl = 0
 
     while teste%2 == 0 :
         teste = teste / 2
         ctrl += 1
 
     return ctrl , teste
 
 
while nqtde > i :
    num1, base = input()
    LT = []
    tmp = 0
    elev =0 
    ctrl, teste = header(num1)
    print ctrl,teste    
 
    potMod(num1, base, teste,1)
 
 
    for k in range(ctrl):
        print(2**k)
        LT.append(2**k)
 
 
    for y in LT:
        tmp += 1 
        exp = y*teste
        j = potMod(num1,base,exp,0)
        print base,exp, num1, j
 
        if ((y == 1) and (j == 1) ) or (y >=1 and j == num1 -1 ) :
            print "INCONCLUSIVO"
            break
 
        if tmp  == len(LT):
            print 'COMPOSTO'
 
 
    print "---"
    i +=1
