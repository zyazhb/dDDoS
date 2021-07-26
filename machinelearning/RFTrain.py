#导入必要的包
from sklearn.ensemble import RandomForestClassifier 
from sklearn.datasets import load_digits
from sklearn.model_selection import train_test_split,GridSearchCV,cross_val_score
from sklearn.metrics import accuracy_score
import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

#导入数据集
data = pd.read_csv('02-14-2018.csv')
# print(data)

# 处理时间数据
import time
timestamp=data.iloc[:,2]
j=0
for i in timestamp:
    # 14/02/2018 08:31:01
    local=time.strptime(i,'%d/%m/%Y %H:%M:%S')
    data.iloc[j,2]=time.mktime(local)
    j+=1
# print(data.iloc[:,2])

x,y=data.iloc[:,:78].values,data.iloc[:,79]
RF = RandomForestClassifier(random_state = 66)
score = cross_val_score(RF,x,y,cv=10).mean()
print('交叉验证得分: %.4f'%score)