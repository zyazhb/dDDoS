# -*- coding: utf-8 -*-
#导入必要的包
from sklearn.ensemble import RandomForestClassifier 
from sklearn.ensemble import ExtraTreesClassifier
from sklearn.model_selection import train_test_split,GridSearchCV,cross_val_score
from sklearn.metrics import accuracy_score
import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
import time

# 处理时间数据
def time_process(data):
    timestamp=data.iloc[:,2]
    j=0
    for i in timestamp:
        # 14/02/2018 08:31:01
        local=time.strptime(i,'%d/%m/%Y %H:%M:%S')
        data.iloc[j,2]=time.mktime(local)
        j+=1
    return data

## 未调参情况
def defaultTest(x,y,newdata):
    clf = RandomForestClassifier(random_state = 66)
    score = cross_val_score(clf,x,y,cv=10).mean()
    print('未调参交叉验证得分: %.6f'%score)
    clf = clf.fit(x,y)
    print('未调参时预测结果为：'+clf.predict(newdata))

## 调参后随机森林
def paramsAdjustTest(x,y,newdata):
    clf2 = RandomForestClassifier(n_estimators=111, max_features=17,random_state=66)
    scores2 = cross_val_score(clf2, x, y)
    print('调参后交叉验证得分: %.6f'%scores2.mean())
    clf2 = clf2.fit(x,y)
    print('调参后预测结果为：'+clf2.predict(newdata))

## ExtraTree分类器集合
def ExtraTreeTest(x,y):
    clf3 = ExtraTreesClassifier(n_estimators=10, max_depth=None,min_samples_split=2, random_state=0)
    scores3 = cross_val_score(clf3, x, y)
    print('ExtraTree分类器集合得分: %.6f'%scores3.mean())

## 参数调整测试
def paramsAdjust():
    param_grid = {'max_features':np.arange(1,80,1)} 
    # param_grid = {'criterion':['gini', 'entropy']}
    rfc = RandomForestClassifier(n_estimators=111
                                ,random_state=0
                                #,max_depth=4  #如果我们加入了map_depth=4分数会更低,可以先加了再注释比较一下分数
                                )
    GS = GridSearchCV(rfc,param_grid,cv=10)
    GS.fit(x,y)
    
    print(GS.best_params_)  #显示调整出来的最佳参数
    print(GS.best_score_)    #分数不变，说明max_depth不影响模型
    #调整max_depth
    # {'max_depth': 1}
    # 0.9909219797216016

    #调整max_features
    # {'max_features': 17}
    # 0.9999507631708517

    # 调整min_samples_leaf
    # {'min_samples_leaf': 1}
    # 0.9909219797216016

    # 调整min_samples_split
    # {'min_samples_split': 2}
    # 0.9998030526834072

    # 调整min_samples_split
    # {'criterion': 'gini'}
    # 0.9909219797216016

## 画出n_estimators的学习曲线，找出最优的n_estimaters
# 0.9998522895125553 111
def find_best_n_estimaters():
    scorel = []    #定义一个列表，用来存放每次循环得到的score
    for i in range(0,200,10):
        rfc = RandomForestClassifier(n_estimators=i+1,  #因为n_estimators不能为0，所以i要加1
                                #   n_jobs=-1,
                                    random_state=0)
        score = cross_val_score(rfc,x,y,cv=10).mean()
        scorel.append(score)
    print(max(scorel),(scorel.index(max(scorel))*10)+1)  #打印出最好的score和他的索引
    plt.figure(figsize=[20,5])
    plt.plot(range(0,200,10),scorel)
    plt.show()
    ## 进一步更小区间的测试
    # scorel = []
    # for i in range(100,125):  
    #     rfc = RandomForestClassifier(n_estimators=i,  #这里从25开始所以不用加1
    #                                 # n_jobs=-1,
    #                                  random_state=0)
    #     score = cross_val_score(rfc,x,y,cv=10).mean()
    #     scorel.append(score)
    # print(max(scorel),([*range(100,125)][scorel.index(max(scorel))]))
    # plt.figure(figsize=[20,5])
    # plt.plot(range(100,125),scorel)
    # plt.show()


if __name__=='__main__':
    #导入数据集
    data = pd.read_csv('test.csv')
    #数据集处理
    time_process(data)
    #设置自变量及因变量
    x,y=data.iloc[:,:79].values,data.iloc[:,79]
    #设置预测数据集
    newdata = [[21,6,1518581347.0,4,1,1,0,0,0,0,0,0,0,0,0,0,0,500000,4,0,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,40,20,250000,250000,0,0,0,0,0,0,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,1,0,26883,0,0,40,0,0,0,0,0,0,0,0]]

    defaultTest(x,y,newdata)
    paramsAdjustTest(x,y,newdata)
    ExtraTreeTest(x,y)