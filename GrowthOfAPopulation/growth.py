import time
import math

def nb_year(p0,percent,aug,p):
    if percent == 0 :
        return (p-p0)/aug
    r = 1+percent/100
    c = (aug+p*r-p)/(aug+p0*r-p0)
    return math.ceil(math.log(c)/math.log(r))

def nb_year_iter(p0,percent,aug,p):
    count = 0
    while p0 < p:
        count += 1
        p0 = (1+percent/100)*p0 + aug
    return count

i = 100
p0 = 100
p = 0.001
aug = 10
pn = 10000000000000
while i < pn:
    start_calc=time.monotonic()
    nb_year(p0,p,aug,i)
    end_calc=time.monotonic()

    start_iter=time.monotonic()
    nb_year_iter(p0,p,aug,i)
    end_iter=time.monotonic()

    print(f"{i} {end_calc-start_calc:g} {end_iter-start_iter:g}")
    i *= 10