from flask import Flask, render_template, request, jsonify

import coeffFunctionsLib
server = Flask(__name__)


@server.route('/coefficient', methods=['POST', 'GET'])
def coefficient():
    form = request.form
    if request.method=='POST':
        formData=request.form #json входные данные c формы
        print(request.form['name'])
        name = formData['name']
        borrCap = int(formData['borrCap'])  # Заемный капитал
        ownCap = int(formData['ownCap'])  # Собственный капитал
        balanceCurr = int(formData['balanceCurr'])  # Валюта баланса
        allCash = int(formData['allCash'])  # Денежные средства
        longTermDuties = int(formData['longTermDuties'])  # Долгосрочные обязательства
        shortTermDuties = int(formData['shortTermDuties'])  # Кратксрочные обязательства
        shortFinInv = int(formData['shortFinInv'])  # Краткосрочные финансовые вложения
        shortRec = int(formData['shortRec'])  # Краткосрочная дебиторская задолженности
        sumMoney = int(formData['sumMoney'])  # Общая сумма ликвидных оборотных средств
        kk = round(coeffFunctionsLib.Kk(int(borrCap), int(ownCap)), 2)
        kn = round(coeffFunctionsLib.Kn(int(ownCap), int(balanceCurr)), 2)
        kfin = round(coeffFunctionsLib.Kfin(int(ownCap), int(borrCap)), 2)
        kfu = round(coeffFunctionsLib.Kfu(int(ownCap), int(longTermDuties), int(balanceCurr)), 2)
        kabsl = round(coeffFunctionsLib.KAbsL(int(allCash), int(shortFinInv), int(shortTermDuties)), 2)
        kfastl = round(coeffFunctionsLib.KFastL(int(allCash), int(shortFinInv), int(shortRec), int(shortTermDuties)), 2)
        kcurrl = round(coeffFunctionsLib.KCurrL(int(sumMoney), int(shortTermDuties)), 2)
        report = coeffFunctionsLib.KReport(kk, kn, kfin, kfu, kabsl, kfastl, kcurrl)
        data = {'kk': kk, 'kn': kn, 'kfin': kfin, 'kfu': kfu, 'kabsl': kabsl, 'kfastl': kfastl, 'kcurrl': kcurrl, 'report': report} #json выходные данные
        return jsonify(data)
        #return render_template('coefficient.html', kk="Коэффициент капитализации: "+str(kk), kn="Коэффициент финансовой независимости: "+str(kn), kfin="Коэффициент финансирования: "+str(kfin), kfu="Коэффициент финанcовой устойчивости: "+str(kfu), kabsl="Коэффициент абсолютной ликвидности: "+str(kabsl), kfastl="Коэффициент быстрой ликвидности: "+str(kfastl), kcurrl="Коэффициент текущей (общей) ликвидности: "+str(kcurrl), titleReport="Отчет", report=report)
    else:
        return render_template('coefficient.html', form=form)

if __name__ == "__main__":
   server.run(host='0.0.0.0')