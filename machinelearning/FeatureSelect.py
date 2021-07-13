import numpy as np
import pandas as pd
from scipy.sparse import data
from sklearn.ensemble import RandomForestClassifier

data=pd.read_csv(r'test.csv')

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

#处理标签数据
# labelname=data.iloc[:,79].unique() #['Benign','FTP-BruteForce']
labelname=data.iloc[:,79]
for i in range(99):
    if((data.iloc[i:,79]=='Benign').item):
        data.iloc[i:,79]=0
    if((data.iloc[i:,79]=='FTP-BruteForce').item):
        data.iloc[i:,79]=1

from sklearn.model_selection import train_test_split
# print(data.shape[1]) #一共80列
x,y=data.iloc[:,:].values,data.iloc[:,78]
x_train,x_test,y_train,y_test = train_test_split(x,y,test_size=0.3,random_state=0)
feat_labels = data.columns

# n_estimators：森林中树的数量
# n_jobs  整数 可选（默认=1） 适合和预测并行运行的作业数，如果为-1，则将作业数设置为核心数
forest = RandomForestClassifier(n_estimators=10000, random_state=0, n_jobs=-1)
forest.fit(x_train, y_train.astype('int'))

# 下面对训练好的随机森林，完成重要性评估
# feature_importances_  可以调取关于特征重要程度
importances = forest.feature_importances_
print("重要性：",importances)
x_columns = data.columns
indices = np.argsort(importances)[::-1]
sort_column=[]
for f in range(x_train.shape[1]):
# 对于最后需要逆序排序，我认为是做了类似决策树回溯的取值，从叶子收敛
# 到根，根部重要程度高于叶子。
    sort_column.append(feat_labels[indices[f]])
    print("%2d) %-*s %f" % (f + 1, 30, feat_labels[indices[f]], importances[indices[f]]))
# 筛选变量（选择重要性比较高的变量）
threshold = 0.15
x_selected = x_train[:,importances > threshold]

label_importance=[]
for k in range(80):
    for d in range(80):
        if(feat_labels[k]==sort_column[d]):
            label_importance.append(79-d)
print(label_importance)
# 可视化
import matplotlib.pyplot as plt
plt.figure(figsize=(10,6))
plt.title("DDoS数据集中各个特征的重要程度",fontsize = 18)
plt.ylabel("import level",fontsize = 15,rotation=90)
plt.rcParams['font.sans-serif'] = ["SimHei"]
plt.rcParams['axes.unicode_minus'] = False
for i in range(x_columns.shape[0]):
    plt.bar(i,importances[indices[i]],color='orange',align='center')
    plt.xticks(np.arange(x_columns.shape[0]),sort_column,rotation=70,fontsize=7)
plt.show()